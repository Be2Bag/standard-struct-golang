package pipeline

import "go.mongodb.org/mongo-driver/bson"

func ExamplePipeline() bson.M {
	return bson.M{
		"$project": bson.M{
			"_id": 1,
		},
	}
}
