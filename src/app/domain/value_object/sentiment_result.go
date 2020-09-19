package value_object

type SentimentResult struct {
	Id    uint64
	Text  string
	Score float32
	Type  SentimentType
}

func NewSentimentResult(
	id uint64,
	text string,
	score float32,
	typ SentimentType,
) *SentimentResult {
	sr := new(SentimentResult)
	sr.Id = id
	sr.Text = text
	sr.Score = score
	sr.Type = typ
	return sr
}
