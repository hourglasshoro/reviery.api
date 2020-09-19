package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/query_service"
)

type GetTwitterAccessTokenQueryService struct {
	Redis *redis.Client
}

func NewGetTwitterAccessTokenQueryService(redis *redis.Client) *GetTwitterAccessTokenQueryService {
	qs := new(GetTwitterAccessTokenQueryService)
	qs.Redis = redis
	return qs
}

func (qs *GetTwitterAccessTokenQueryService) GetTwitterToken() (token value_object.AccessToken, err error) {
	tokens, err := qs.Redis.Keys(ctx, "twitterToken:*").Result()
	if len(tokens) == 0 {
		err = query_service.NoTwitterAccessTokenExistException
	}
	if err != nil {
		return
	}
	tokString, err := qs.Redis.Get(ctx, tokens[0]).Result()
	if err != nil {
		return
	}
	tok, err := value_object.NewAccessToken(tokString)
	if err != nil {
		return
	}
	token = *tok
	return
}
