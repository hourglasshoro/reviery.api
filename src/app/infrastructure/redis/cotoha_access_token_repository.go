package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	"time"
)

type CotohaAccessTokenRepository struct {
	Redis *redis.Client
}

func NewCotohaAccessTokenRepository(
	redis *redis.Client,
) *CotohaAccessTokenRepository {
	repo := new(CotohaAccessTokenRepository)
	repo.Redis = redis
	return repo
}

func (repo *CotohaAccessTokenRepository) SetCotohaToken(token value_object.AccessToken, ttl time.Duration) (err error) {
	key := fmt.Sprintf("CotohaToken:%s", uuid.New())
	err = repo.Redis.Set(ctx, key, token.Token, ttl).Err()
	return
}
