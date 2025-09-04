package example_repo

import (
	"context"
	"errors"
	"fmt"
	"standard-struct-golang/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrExampleNotFound = errors.New("example not found")

func cloneAndInjectNotDeleted(filter interface{}) interface{} {
	m, ok := filter.(bson.M)
	if !ok {
		return filter
	}
	cp := make(bson.M, len(m)+1)
	for k, v := range m {
		cp[k] = v
	}
	if _, exists := cp["deleted_at"]; !exists {
		if _, hasOr := cp["$or"]; !hasOr {
			cp["$or"] = []bson.M{{"deleted_at": bson.M{"$exists": false}}, {"deleted_at": nil}}
		}
	}
	return cp
}

func (r *ExampleRepositories) CreateExample(ctx context.Context, example models.Example) (*models.Example, error) {
	now := time.Now()
	if example.CreatedAt.IsZero() {
		example.CreatedAt = now
	}
	example.UpdatedAt = now

	_, err := r.DB().Collection(example.CollectionName()).InsertOne(ctx, example)
	if err != nil {
		return nil, fmt.Errorf("insert example: %w", err)
	}
	return &example, nil
}

func (r *ExampleRepositories) UpdateExampleByID(ctx context.Context, exampleID string, example models.Example) (*models.Example, error) {
	if exampleID == "" {
		return nil, errors.New("exampleID is required")
	}
	col := r.DB().Collection(models.Example{}.CollectionName())

	filter := bson.M{"example_id": exampleID, "$or": []bson.M{{"deleted_at": bson.M{"$exists": false}}, {"deleted_at": nil}}}

	set := bson.M{
		"updated_at": time.Now(),
	}
	if example.Detail != "" {
		set["detail"] = example.Detail
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated models.Example
	if err := col.FindOneAndUpdate(ctx, filter, bson.M{"$set": set}, opts).Decode(&updated); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrExampleNotFound
		}
		return nil, fmt.Errorf("update example by id: %w", err)
	}
	return &updated, nil
}

func (r *ExampleRepositories) SoftDeleteExampleByID(ctx context.Context, exampleID string) error {
	if exampleID == "" {
		return errors.New("exampleID is required")
	}
	col := r.DB().Collection(models.Example{}.CollectionName())
	now := time.Now()
	res, err := col.UpdateOne(ctx,
		bson.M{"example_id": exampleID, "$or": []bson.M{{"deleted_at": bson.M{"$exists": false}}, {"deleted_at": nil}}},
		bson.M{"$set": bson.M{"deleted_at": now, "updated_at": now}},
	)
	if err != nil {
		return fmt.Errorf("soft delete example: %w", err)
	}
	if res.MatchedCount == 0 {
		return ErrExampleNotFound
	}
	return nil
}

func (r *ExampleRepositories) FindExamples(ctx context.Context, filter interface{}, projection interface{}) ([]*models.Example, error) {
	filter = cloneAndInjectNotDeleted(filter)

	opts := options.Find()
	if projection != nil {
		opts.SetProjection(projection)
	}

	cur, err := r.DB().Collection(models.Example{}.CollectionName()).Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("find examples: %w", err)
	}
	defer cur.Close(ctx)

	var examples []*models.Example
	for cur.Next(ctx) {
		var ex models.Example
		if err := cur.Decode(&ex); err != nil {
			return nil, fmt.Errorf("decode example: %w", err)
		}
		examples = append(examples, &ex)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("cursor err: %w", err)
	}
	return examples, nil
}

func (r *ExampleRepositories) FindExample(ctx context.Context, filter interface{}, projection interface{}) (*models.Example, error) {
	filter = cloneAndInjectNotDeleted(filter)

	opts := options.FindOne()
	if projection != nil {
		opts.SetProjection(projection)
	}
	var example models.Example
	if err := r.DB().Collection(example.CollectionName()).FindOne(ctx, filter, opts).Decode(&example); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrExampleNotFound
		}
		return nil, fmt.Errorf("find example: %w", err)
	}
	return &example, nil
}

func (r *ExampleRepositories) FindExamplesPaged(ctx context.Context, filter interface{}, projection interface{}, sort bson.D, skip, limit int64) ([]models.Example, int64, error) {
	filter = cloneAndInjectNotDeleted(filter)

	findOpts := options.Find().
		SetSort(sort).
		SetSkip(skip).
		SetLimit(limit)
	if projection != nil {
		findOpts.SetProjection(projection)
	}

	col := r.DB().Collection(models.Example{}.CollectionName())

	cur, err := col.Find(ctx, filter, findOpts)
	if err != nil {
		return nil, 0, fmt.Errorf("find paged examples: %w", err)
	}
	defer cur.Close(ctx)

	var results []models.Example
	if err := cur.All(ctx, &results); err != nil {
		return nil, 0, fmt.Errorf("decode paged examples: %w", err)
	}

	total, err := col.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, fmt.Errorf("count examples: %w", err)
	}

	return results, total, nil
}
