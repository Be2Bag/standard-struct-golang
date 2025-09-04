package services

import (
	"context"
	"fmt"
	"standard-struct-golang/modules/frontweb/middleware"
	"standard-struct-golang/packages/health_id"
	"standard-struct-golang/packages/provider"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func (svc AuthService) LoginHealthId(ctx context.Context, code string, uri string) (*health_id.ResponseHealthIdToken, int, error) {
	ctx, span := otel.Tracer("frontweb").Start(ctx, "LoginHealthId", oteltrace.WithAttributes())
	defer span.End()

	auth, _, statusCode, _, _, err := svc.repo.HealthId().GetHealthIdTokenByCode(ctx, code, uri)
	if err != nil {
		return nil, statusCode, err
	}
	return auth, statusCode, nil
}

func (svc AuthService) GetProviderToken(ctx context.Context, accessToken string) (*provider.ResponseProviderToken, int, error) {
	ctx, span := otel.Tracer("frontweb").Start(ctx, "GetProviderToken", oteltrace.WithAttributes())
	defer span.End()

	provider, _, statusCode, _, _, err := svc.repo.Provider().GetProviderToken(ctx, accessToken)
	if err != nil {
		return nil, statusCode, err
	}
	return provider, statusCode, nil

}

func (svc AuthService) GetProviderData(ctx context.Context, accessToken string) (*provider.ProviderData, int, error) {
	ctx, span := otel.Tracer("frontweb").Start(ctx, "GetProviderData", oteltrace.WithAttributes())
	defer span.End()

	providerData, _, statusCode, _, _, err := svc.repo.Provider().GetProviderData(ctx, accessToken)
	if err != nil {
		return nil, statusCode, err
	}

	return providerData, statusCode, nil
}

func (svc AuthService) CreateSession(ctx context.Context, provider *provider.ProviderData, providerAccessToken string) (string, error) {
	ctx, span := otel.Tracer("frontweb").Start(ctx, "CreateSession", oteltrace.WithAttributes())
	defer span.End()

	sessionID := uuid.New().String()

	sessionData := middleware.UserSessionsClaims{
		HashCID:       provider.HashCid,
		ProviderToken: providerAccessToken,
		ProviderID:    provider.ProviderID,
	}
	if err := svc.repo.Cache().StoredRegisterSession(ctx, sessionID, sessionData); err != nil {
		return "", fmt.Errorf("error storing session data: %w", err)
	}

	claims := &middleware.Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   sessionID,
			Audience:  provider.ProviderID,
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "standard-struct-golang",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(svc.repo.AppCfg().CredentialConfig.AuthEncryptRegisterKey))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}
