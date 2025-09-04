package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"standard-struct-golang/appconst"
	dtoDefault "standard-struct-golang/modules/frontweb/modules/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type SessionStore interface {
	GetAuthSession(key string) (UserSessionsClaims, error)
}

type AuthMiddleware struct {
	AuthEncryptKey *string
	SessionStore   SessionStore
	OpenApiSecret  *string
}

func NewAuthMiddleware(authEncryptKey *string, sessionStore SessionStore, openApiSecret *string) *AuthMiddleware {
	return &AuthMiddleware{
		AuthEncryptKey: authEncryptKey,
		SessionStore:   sessionStore,
		OpenApiSecret:  openApiSecret,
	}
}

func (am AuthMiddleware) NewSessionMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Cookies(appconst.AuthService)
		if tokenString == "" {
			return rejectUnauthorizedRequest(ctx, "no cookie attached", "ไม่พบคุกกี้")
		}

		sessionID, err := am.verifySessionToken(tokenString)
		if err != nil {
			log.Printf("Token verification failed: %v", err)
			return rejectUnauthorizedRequest(ctx, "invalid token: "+err.Error(), "token ไม่ถูกต้อง")
		}

		sessionData, err := am.SessionStore.GetAuthSession(sessionID)
		if err != nil {
			return rejectUnauthorizedRequest(ctx, "invalid or expired session", "Session ไม่ถูกต้องหรือหมดอายุ")
		}

		currentUA := ctx.Get("User-Agent")
		if sessionData.UserAgent != currentUA {
			return ctx.Status(fiber.StatusUnauthorized).JSON(dtoDefault.BaseResponse{
				StatusCode: fiber.StatusUnauthorized,
				MessageEN:  "Session User-Agent mismatch",
				MessageTH:  "User-Agent ของ Session ไม่ตรงกัน",
				Status:     fiber.ErrUnauthorized.Message,
				Data:       nil,
			})
		}

		AuthService := context.WithValue(ctx.Context(), appconst.AuthService, &sessionData)
		ctx.SetUserContext(AuthService)
		return ctx.Next()
	}
}

func (am AuthMiddleware) verifySessionToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(*am.AuthEncryptKey), nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	if time.Now().Unix() > claims.StandardClaims.ExpiresAt {
		return "", errors.New("token expired")
	}

	sessionID := claims.StandardClaims.Subject
	if sessionID == "" {
		return "", errors.New("missing session ID in token")
	}
	return sessionID, nil
}

func rejectUnauthorizedRequest(ctx *fiber.Ctx, errMsgEN, errMsgTH string) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(dtoDefault.BaseResponse{
		StatusCode: fiber.StatusUnauthorized,
		MessageEN:  errMsgEN,
		MessageTH:  errMsgTH,
		Status:     fiber.ErrUnauthorized.Message,
		Data:       nil,
	})
}

func (am AuthMiddleware) GetVerifiedClaims(ctx *fiber.Ctx) (*UserSessionsClaims, error) {
	tokenString := ctx.Cookies(appconst.AuthService)
	if tokenString == "" {
		return nil, errors.New("no cookie attached")
	}

	sessionID, err := am.verifySessionToken(tokenString)
	if err != nil {
		return nil, err
	}

	sessionData, err := am.SessionStore.GetAuthSession(sessionID)
	if err != nil {
		return nil, err
	}

	return &sessionData, nil
}
