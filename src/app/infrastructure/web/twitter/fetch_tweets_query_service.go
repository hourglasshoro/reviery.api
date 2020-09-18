package twitter

import (
	"encoding/json"
	"fmt"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/query_service"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/query_service/read_model"
	"net/http"
	"strconv"
)

const MaxTweetsResultNumber = 10

type FetchTweetsQueryService struct {
}

type TweetsResponse struct {
	Data []struct {
		Id   int    `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
}

func (qs *FetchTweetsQueryService) FetchTweets(keyword value_object.Keyword, token value_object.AccessToken) (readModel read_model.FetchTweetsReadModel, err error) {
	url := fmt.Sprintf("https://api.twitter.com/2/tweets/search/recent?max_results=%d&query=%s", MaxTweetsResultNumber, keyword)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return
	}
	if res.Status != strconv.Itoa(http.StatusOK) {
		err = query_service.CannotFetchTweetsException
	}
	defer res.Body.Close()
	var tweetsRes TweetsResponse
	err = json.NewDecoder(res.Body).Decode(&tweetsRes)
	if err != nil {
		return
	}
	var tweets []read_model.TweetReadModel
	for _, v := range tweetsRes.Data {
		tweets = append(tweets, read_model.TweetReadModel{
			Id:   uint64(v.Id),
			Text: v.Text,
		})
	}
	readModel = read_model.FetchTweetsReadModel{Tweets: tweets}
	return
}
