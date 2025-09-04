package example_port

import (
	"context"
	"standard-struct-golang/config"
	"standard-struct-golang/models"
	"standard-struct-golang/packages/cache/cache"
	"standard-struct-golang/packages/health_id"
	"standard-struct-golang/packages/moph_account_center"
	"standard-struct-golang/packages/moph_line"
	"standard-struct-golang/packages/provider"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/trace"
)

type ExampleRepositories interface {
	Module() string
	AppCfg() *config.Config
	Log() *logrus.Entry
	DB() *mongo.Database
	MophAc() *moph_account_center.Client
	MophLine() *moph_line.Client
	Trace(ctx context.Context, spanName string, attributes ...trace.SpanStartOption) (context.Context, trace.Span)
	HealthId() *health_id.HealthId
	Provider() *provider.Provider
	Cache() *cache.Cache

	CreateExample(ctx context.Context, example models.Example) (*models.Example, error)
	UpdateExampleByID(ctx context.Context, exampleID string, example models.Example) (*models.Example, error)
	SoftDeleteExampleByID(ctx context.Context, exampleID string) error
	FindExamples(ctx context.Context, filter interface{}, projection interface{}) ([]*models.Example, error)
	FindExample(ctx context.Context, filter interface{}, projection interface{}) (*models.Example, error)
	FindExamplesPaged(ctx context.Context, filter interface{}, projection interface{}, sort bson.D, skip, limit int64) ([]models.Example, int64, error)
}

type ExampleService interface {
	CreateExample(ctx context.Context, example models.Example) error
	GetExampleByID(ctx context.Context, exampleID string) (*models.Example, error)
	SendMophLineNotify(ctx context.Context, cid string) error
}
