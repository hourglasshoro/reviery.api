package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/query_service"
)

type GetCotohaAccessTokenQueryService struct {
	Redis *redis.Client
}

func (qs *GetCotohaAccessTokenQueryService) GetCotohaToken() (token value_object.AccessToken, err error) {
	tokens, err := qs.Redis.Keys(ctx, "CotohaToken:*").Result()
	if len(tokens) == 0 {
		err = query_service.NoCotohaAccessTokenExistException
	}
	if err != nil {
		return
	}
	tok, err := value_object.NewAccessToken(tokens[0])
	if err != nil {
		return
	}
	token = *tok
	return
}
