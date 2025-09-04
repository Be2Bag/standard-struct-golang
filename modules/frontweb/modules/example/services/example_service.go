package services

import (
	"context"
	"errors"
	"standard-struct-golang/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.opentelemetry.io/otel/attribute"

	oteltrace "go.opentelemetry.io/otel/trace"
)

func (svc ExampleService) CreateExample(ctx context.Context, example models.Example) error {

	if example.ExampleID == "" {
		return errors.New("exampleID is required")
	}

	if example.Detail == "" {
		return errors.New("detail is required")
	}

	_, errOnCreateExample := svc.repo.CreateExample(ctx, example)
	if errOnCreateExample != nil {
		return errOnCreateExample
	}

	return nil
}

func (svc ExampleService) GetExampleByID(ctx context.Context, exampleID string) (*models.Example, error) {

	if exampleID == "" {
		return nil, errors.New("exampleID is required")
	}

	filter := bson.M{"example_id": exampleID}
	projection := bson.M{}

	example, err := svc.repo.FindExample(ctx, filter, projection)
	if err != nil {
		return nil, err
	}

	return example, nil
}

func (svc ExampleService) SendMophLineNotify(ctx context.Context, cid string) error {

	ctx, span := svc.repo.Trace(ctx, "svc.SendMophLineNotify", oteltrace.WithAttributes(
		attribute.String("cid", cid),
	))
	defer span.End()

	var token *string
	var err error

	token, err = svc.repo.Cache().GetLineSession(ctx, "line_access_token")
	if err != nil || token == nil || *token == "" {
		tokenNew := svc.repo.MophAc().GetToken(ctx)
		svc.repo.Cache().StoredLineSession(ctx, "line_access_token", tokenNew)
		token = &tokenNew
	}

	_, _, _, _, errOnSendLine := svc.repo.MophLine().SendLine(ctx, *token, cid)
	if errOnSendLine != nil {
		return errOnSendLine
	}

	return nil

}
