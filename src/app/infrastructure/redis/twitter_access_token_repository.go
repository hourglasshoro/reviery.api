package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	"time"
)

type TwitterAccessTokenRepository struct {
	Redis *redis.Client
	TTL   time.Duration
}

func NewTwitterAccessTokenRepository(
	redis *redis.Client,
	ttl time.Duration,
) *TwitterAccessTokenRepository {
	repo := new(TwitterAccessTokenRepository)
	repo.Redis = redis
	repo.TTL = ttl
	return repo
}

func (repo *TwitterAccessTokenRepository) SetTwitterToken(token value_object.AccessToken) (err error) {
	key := fmt.Sprintf("twitterToken:%s", uuid.New())
	err = repo.Redis.Set(ctx, key, token, repo.TTL).Err()
	return
}
