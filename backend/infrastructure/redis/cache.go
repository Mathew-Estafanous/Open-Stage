package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

type redisMemoryCache struct {
	client *redis.Client
}

func NewMemoryCache(client *redis.Client) *redisMemoryCache {
	return &redisMemoryCache{
		client: client,
	}
}

func (r redisMemoryCache) Store(tkn string) error {
	parser := jwt.Parser{
		UseJSONNumber: true,
	}
	tk, err := parser.Parse(tkn, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return fmt.Errorf("%w: issue while parsing token", domain.Internal)
	}

	jsonExp := tk.Claims.(jwt.MapClaims)["exp"].(json.Number)
	exp, err := jsonExp.Int64()
	if err != nil {
		return fmt.Errorf("%w: error while parsing json number", domain.Internal)
	}
	dif := exp  - time.Now().Unix()
	dur := time.Duration(dif) * time.Second

	ctx := context.Background()
	_, err = r.client.SetEX(ctx, tkn, "", dur).Result()
	return err
}

func (r redisMemoryCache) Contains(tkn string) (bool, error) {
	ctx := context.Background()
	res, err := r.client.Exists(ctx, tkn).Result()
	if err != nil {
		return false, fmt.Errorf("%w: error while checking if token cached", err)
	}

	return res == 1, nil
}