package read_model

type FetchTweetsReadModel struct {
	Tweets []TweetReadModel
}

type TweetReadModel struct {
	Id   uint64
	Text string
}
