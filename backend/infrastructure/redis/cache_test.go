package redis

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestRedisMemoryCache_StoreToken(t *testing.T) {
	err := os.Setenv("SECRET_KEY", "SECRET")
	assert.NoError(t, err)

	exp := time.Now().Add(time.Minute * 5).Unix()
	refreshClaim := jwt.StandardClaims{
			ExpiresAt: exp,
			Audience:  "access",
			Subject:   "2",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	jwtTkn, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	assert.NoError(t, err)

	client, mock := redismock.NewClientMock()
	cache := NewMemoryCache(client)

	mock.ExpectSetEX(jwtTkn, "", time.Minute * 5).SetVal("OK")

	err = cache.Store(jwtTkn)
	assert.NoError(t, err)
}

func TestRedisMemoryCache_ContainsToken(t *testing.T) {
	client, mock := redismock.NewClientMock()
	cache := NewMemoryCache(client)

	mock.ExpectExists("SomeJWT").SetVal(1)

	res, err := cache.Contains("SomeJWT")
	assert.NoError(t, err)
	assert.Equal(t, true, res)

	mock.ExpectExists("NotInCacheJWT").SetVal(0)

	res, err = cache.Contains("NotInCacheJWT")
	assert.NoError(t, err)
	assert.Equal(t, false, res)
}
