package query_service

import (
	"errors"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/query_service/read_model"
)

var CannotFetchTweetsException = errors.New("cannot fetch tweets")

type FetchTweetsQueryService interface {
	FetchTweets(keyword value_object.Keyword, token value_object.AccessToken) (readModel read_model.FetchTweetsReadModel, err error)
}
