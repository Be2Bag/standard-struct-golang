package models

import "time"

type Example struct {
	ExampleID string     `bson:"example_id"`
	Detail    string     `bson:"detail"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at"`
}

func (Example) CollectionName() string {
	return "examples"
}
