package value_object

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
