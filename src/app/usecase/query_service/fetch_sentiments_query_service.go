package query_service

import (
	"github.com/hourglasshoro/reviery.api/src/app/domain/entitiy"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/query_service/read_model"
)

type FetchSentimentQueryService interface {
	FetchSentiment(opinions []entitiy.Opinion, token value_object.AccessToken) (readModel read_model.FetchSentimentsReadModel, err error)
}
