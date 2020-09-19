package read_model

import "github.com/hourglasshoro/reviery.api/src/app/domain/value_object"

type FetchSentimentsReadModel struct {
	List []SentimentReadModel
}

type SentimentReadModel struct {
	Id    uint64
	Text  string
	Score float32
	Type  value_object.SentimentType
}
