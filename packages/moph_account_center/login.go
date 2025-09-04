package moph_account_center

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/bytedance/sonic"
)

type Credential struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	mutex     *sync.Mutex
}

type JwtTokenClaims struct {
	Exp int `json:"exp"`
}

func (c *Client) Login(ctx context.Context) error {
	ctx, span := c.tracer.Start(ctx, "moph_account_center.login")
	defer span.End()
	url := fmt.Sprintf("%s/token?Action=get_moph_access_token&hospital_code=%s&password_hash=%s&user=%s", c.url, c.hCode, c.hashPassword, c.username)
	resp, err := c.http.Post(ctx, url, nil, nil, 10)

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("request time out after 30 seconds")
		}
		return fmt.Errorf("failed to make POST request: %w", err)
	}

	if resp.Code == http.StatusOK {
		token := string(resp.Body)
		c.credential.mutex.Lock()
		c.credential.Token = token
		c.credential.ExpiresAt = time.Now().Add(time.Hour * 3)
		expiresTime, err := c.GetExpiresTime(token)
		if err == nil {
			c.credential.ExpiresAt = *expiresTime
		}
		c.credential.mutex.Unlock()
		return nil
	}
	return errors.New("Service Unavailable")
}

func (c *Client) GetToken(ctx context.Context) string {
	ctx, span := c.tracer.Start(ctx, "moph_account_center.getToken")
	defer span.End()

	if c.credential.Token == "" || c.credential.ExpiresAt.Before(time.Now()) {
		if err := c.Login(ctx); err != nil {
			return ""
		}
	}
	return c.credential.Token
}

func (c *Client) GetExpiresTime(t string) (*time.Time, error) {
	jwtBody := strings.Split(t, ".")
	if len(jwtBody) != 3 {
		return nil, errors.New("invalid jwt token")
	}
	b64Body, err := base64.RawURLEncoding.DecodeString(jwtBody[1])
	if err != nil {
		return nil, err
	}
	var claims JwtTokenClaims
	if err := sonic.Unmarshal(b64Body, &claims); err != nil {
		return nil, err
	}
	expTime := time.Unix(int64(claims.Exp), 0)
	return &expTime, nil
}
