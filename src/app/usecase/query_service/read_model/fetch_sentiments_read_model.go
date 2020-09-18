package read_model

type FetchSentimentsReadModel struct {
	List []SentimentReadModel
}

type SentimentReadModel struct {
	Id    uint64
	Text  string
	Score float32
	Type  SentimentType
}

type SentimentType struct {
	Type string
}

var SentimentTypes = struct {
	Positive SentimentType
	Negative SentimentType
	Neutral  SentimentType
}{
	Positive: SentimentType{"Positive"},
	Negative: SentimentType{"Negative"},
	Neutral:  SentimentType{"Neutral"},
}

func (status *SentimentType) String() string {
	return status.Type
}
