package auth_port

import (
	"context"
	"standard-struct-golang/config"
	"standard-struct-golang/packages/cache/cache"
	"standard-struct-golang/packages/health_id"
	"standard-struct-golang/packages/moph_account_center"
	"standard-struct-golang/packages/moph_line"
	"standard-struct-golang/packages/provider"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/trace"
)

type AuthRepositories interface {
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
}

type AuthService interface {
	LoginHealthId(ctx context.Context, code string, uri string) (*health_id.ResponseHealthIdToken, int, error)
	GetProviderToken(ctx context.Context, accessToken string) (*provider.ResponseProviderToken, int, error)
	GetProviderData(ctx context.Context, accessToken string) (*provider.ProviderData, int, error)
	CreateSession(ctx context.Context, provider *provider.ProviderData, providerAccessToken string) (string, error)
}
