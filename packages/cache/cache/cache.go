package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"standard-struct-golang/config"
	"standard-struct-golang/modules/frontweb/middleware"
	"standard-struct-golang/packages/cache/keydb"

	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
	config config.CacheConfig
	Keys   *keydb.KeyDB
}

func New(cfg config.CacheConfig) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("failed to connect to Redis cache: %v", err)
	} else {
		log.Println("successfully connected to Redis")
	}

	return &Cache{
		client: client,
		config: cfg,
		Keys:   keydb.New(cfg.KeyPrefix),
	}
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, timeout time.Duration) error {
	j, err := sonic.Marshal(&value)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, j, timeout).Err()
}

func (c *Cache) Get(ctx context.Context, key string, value interface{}) error {
	b, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, value)
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

func (c *Cache) GetLineSession(ctx context.Context, sessionID string) (*string, error) {
	key := c.Keys.LineKey(sessionID)
	var token string
	if err := c.Get(ctx, key, &token); err != nil {
		return &token, errors.New("session not found or expired")
	}
	return &token, nil
}

func (c *Cache) StoredLineSession(ctx context.Context, sessionID string, token string) error {
	key := c.Keys.LineKey(sessionID)
	return c.Set(ctx, key, token, 30)
}

func (c *Cache) StoredAuthSession(ctx context.Context, sessionID string, sessionData middleware.UserSessionsClaims) error {
	key := c.Keys.AuthSessionKey(sessionID)
	return c.Set(ctx, key, sessionData, time.Hour*24)
}

func (c *Cache) GetAuthSession(ctx context.Context, sessionID string) (middleware.UserSessionsClaims, error) {
	key := c.Keys.AuthSessionKey(sessionID)
	var sessionData middleware.UserSessionsClaims
	if err := c.Get(ctx, key, &sessionData); err != nil {
		return middleware.UserSessionsClaims{}, errors.New("session not found or expired")
	}
	return sessionData, nil
}

func (c *Cache) InvalidateAuthSession(ctx context.Context, sessionID string) error {
	key := c.Keys.AuthSessionKey(sessionID)
	return c.client.Del(ctx, key).Err()
}

func (c *Cache) StoredRegisterSession(ctx context.Context, sessionID string, sessionData middleware.UserSessionsClaims) error {
	key := c.Keys.RegisterKey(sessionID)
	return c.Set(ctx, key, sessionData, time.Minute*15)
}

func (c *Cache) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}
