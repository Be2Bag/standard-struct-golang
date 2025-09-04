package example_repo

import (
	"context"
	"standard-struct-golang/config"
	repositories "standard-struct-golang/modules/frontweb/repo"
	"standard-struct-golang/packages/cache/cache"
	"standard-struct-golang/packages/health_id"
	"standard-struct-golang/packages/moph_account_center"
	"standard-struct-golang/packages/moph_line"
	"standard-struct-golang/packages/provider"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const moduleName = "example"

type ExampleRepositories struct {
	config        *config.Config
	log           *logrus.Entry
	databaseMongo *mongo.Database
	mophLine      *moph_line.Client
	mophAcClient  *moph_account_center.Client
	tracer        trace.Tracer
	healthId      *health_id.HealthId
	provider      *provider.Provider
	cache         *cache.Cache
}

func NewExampleRepo(frontRepo *repositories.Repository) *ExampleRepositories {

	return &ExampleRepositories{
		config:        frontRepo.AppCfg(),
		log:           frontRepo.Log().WithField("module", moduleName),
		databaseMongo: frontRepo.DB(),
		mophLine:      frontRepo.MophLine(),
		mophAcClient:  frontRepo.MophAc(),
		tracer:        otel.Tracer(moduleName),
		healthId:      frontRepo.HealthId(),
		provider:      frontRepo.Provider(),
		cache:         frontRepo.Cache(),
	}
}

func (r ExampleRepositories) Module() string {
	return moduleName
}

func (r ExampleRepositories) AppCfg() *config.Config {
	return r.config
}

func (r ExampleRepositories) Log() *logrus.Entry {
	return r.log.Dup()
}

func (r ExampleRepositories) DB() *mongo.Database {
	return r.databaseMongo
}

func (r ExampleRepositories) MophLine() *moph_line.Client {
	return r.mophLine
}

func (r ExampleRepositories) MophAc() *moph_account_center.Client {
	return r.mophAcClient
}

func (r ExampleRepositories) Trace(ctx context.Context, spanName string, attributes ...trace.SpanStartOption) (context.Context, trace.Span) {
	return r.tracer.Start(ctx, spanName, attributes...)
}

func (r ExampleRepositories) HealthId() *health_id.HealthId {
	return r.healthId
}

func (r ExampleRepositories) Provider() *provider.Provider {
	return r.provider
}

func (r ExampleRepositories) Cache() *cache.Cache {
	return r.cache
}
