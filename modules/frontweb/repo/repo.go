package repositories

import (
	"context"
	"standard-struct-golang/app"
	"standard-struct-golang/config"
	"standard-struct-golang/packages/cache/cache"
	"standard-struct-golang/packages/health_id"
	"standard-struct-golang/packages/mongodb"
	"standard-struct-golang/packages/moph_account_center"
	"standard-struct-golang/packages/moph_line"
	"standard-struct-golang/packages/provider"
	"standard-struct-golang/packages/requests"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const moduleName = "frontweb"

type Repository struct {
	app           *app.App
	mongodbClient mongodb.Client
	log           *logrus.Entry
	mophLine      *moph_line.Client
	mophAcClient  *moph_account_center.Client
	tracer        trace.Tracer
	healthId      *health_id.HealthId
	provider      *provider.Provider
	cache         *cache.Cache
}

func NewRepo(app *app.App) (*Repository, error) {
	l := app.NewLogger().WithField("module", moduleName)
	httpClient := requests.NewHttpClient(l)

	mongoClient := mongodb.NewWithConnectionString(app.Config.MongoConfig.Connection)
	errorOnConnect := mongoClient.Connect()
	if errorOnConnect != nil {
		return nil, errorOnConnect
	}

	mophLine := moph_line.New(app.Config.MophLine.UrlSendNoti, httpClient, l)
	mophAcClient := moph_account_center.New(app.Config.MophLine.UrlLineLogin, app.Config.MophLine.UserLineNoti, app.Config.MophLine.PasswordHashLineNoti, app.Config.MophLine.HoscodeLineNoti, httpClient, l)
	return &Repository{
		app:           app,
		log:           l,
		mongodbClient: mongoClient,
		mophLine:      mophLine,
		mophAcClient:  mophAcClient,
		tracer:        otel.Tracer(moduleName),
		healthId:      health_id.New(httpClient, app.Config.HealthIdConfig.Url, app.Config.HealthIdConfig.ClientId, app.Config.HealthIdConfig.SecretKey, app.Config.HealthIdConfig.RedirectUrl, app.Config.HealthIdConfig.RedirectLocalhost, app.Config.HealthIdConfig.MophUrl, app.Config.HealthIdConfig.MophClientId, app.Config.HealthIdConfig.MophSecret, l, app.Config.HealthIdConfig.Timeout),
		provider:      provider.New(httpClient, app.Config.ProviderConfig.Url, app.Config.ProviderConfig.Redirect, app.Config.ProviderConfig.ClientId, app.Config.ProviderConfig.SecretKey, l, app.Config.ProviderConfig.Timeout),
		cache:         cache.New(app.Config.CacheConfig),
	}, nil
}

func (r Repository) Module() string {
	return moduleName
}

func (r Repository) AppCfg() *config.Config {
	return r.app.Config
}

func (r Repository) Log() *logrus.Entry {
	return r.log.Dup()
}

func (r Repository) DB() *mongo.Database {
	return r.mongodbClient.GetClient().Database(r.AppCfg().MongoConfig.DatabaseName)
}

func (r Repository) MophLine() *moph_line.Client {
	return r.mophLine
}

func (r Repository) MophAc() *moph_account_center.Client {
	return r.mophAcClient
}

func (r Repository) Trace(ctx context.Context, spanName string, attributes ...trace.SpanStartOption) (context.Context, trace.Span) {
	return r.tracer.Start(ctx, spanName, attributes...)
}

func (r Repository) HealthId() *health_id.HealthId {
	return r.healthId
}

func (r Repository) Provider() *provider.Provider {
	return r.provider
}

func (r Repository) Cache() *cache.Cache {
	return r.cache
}
