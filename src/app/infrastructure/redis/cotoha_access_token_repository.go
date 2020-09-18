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
	TTL   time.Duration
}

func NewCotohaAccessTokenRepository(
	redis *redis.Client,
	ttl time.Duration,
) *CotohaAccessTokenRepository {
	repo := new(CotohaAccessTokenRepository)
	repo.Redis = redis
	repo.TTL = ttl
	return repo
}

func (repo *CotohaAccessTokenRepository) SetCotohaToken(token value_object.AccessToken) (err error) {
	key := fmt.Sprintf("CotohaToken:%s", uuid.New())
	err = repo.Redis.Set(ctx, key, token, repo.TTL).Err()
	return
}
