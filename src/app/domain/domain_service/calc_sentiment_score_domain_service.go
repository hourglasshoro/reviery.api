package domain_service

import (
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/query_service/read_model"
)

const MaxRating = 5

type CalcSentimentScoreDomainService struct {
}

func NewCalcSentimentScoreDomainService() *CalcSentimentScoreDomainService {
	ds := new(CalcSentimentScoreDomainService)
	return ds
}

func (ds *CalcSentimentScoreDomainService) CalcScore(opinions read_model.FetchSentimentsReadModel) (score float32, results []value_object.SentimentResult, err error) {
	total := float32(0)
	for _, v := range opinions.List {
		switch v.Type {
		case value_object.SentimentTypes.Positive:
			cs := convertScore(v.Score)
			results = append(results, *value_object.NewSentimentResult(v.Id, v.Text, cs, v.Type))
			total += cs
		case value_object.SentimentTypes.Negative:
			cs := convertScore(v.Score * -1)
			results = append(results, *value_object.NewSentimentResult(v.Id, v.Text, cs, v.Type))
			total += cs
		case value_object.SentimentTypes.Neutral:
			//cs := convertScore(v.Score)
			//results = append(results, *NewSentimentResult(v.Id, v.Text, cs, v.Type))
			//total += cs
		}
	}
	score = total / float32(len(results))
	return
}

func convertScore(score float32) float32 {
	return (score + 1) * MaxRating / 2
}
