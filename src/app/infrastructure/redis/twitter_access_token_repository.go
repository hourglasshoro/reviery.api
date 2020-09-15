package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/hourglasshoro/reviery.api/src/app/infrastructure/web/twitter"
	"time"
)

type TwitterAccessTokenRepository struct {
	Redis                 *redis.Client
	TTL                   time.Duration
	AuthorizeQueryService twitter.AuthorizeQueryService
}

func NewTwitterAccessTokenRepository(
	redis *redis.Client,
	ttl time.Duration,
	authorizeQueryService twitter.AuthorizeQueryService) *TwitterAccessTokenRepository {
	repo := new(TwitterAccessTokenRepository)
	repo.Redis = redis
	repo.TTL = ttl
	repo.AuthorizeQueryService = authorizeQueryService
	return repo
}

func (repo *TwitterAccessTokenRepository) Set(token string) (err error) {
	key := fmt.Sprintf("token:%s", uuid.New())
	err = repo.Redis.Set(ctx, key, token, repo.TTL).Err()
	return
}

func (repo *TwitterAccessTokenRepository) Get() (token string, err error) {
	tokens, err := repo.Redis.Keys(ctx, "token:*").Result()
	if err != nil {
		return
	}
	if len(tokens) == 0 {
		tok, err := repo.AuthorizeQueryService.Auth()
		if err != nil {
			return
		}
		err = repo.Set(tok)
		if err != nil {
			return
		}
		token = tok
		return
	} else {
		token = tokens[0]
		return
	}
}
