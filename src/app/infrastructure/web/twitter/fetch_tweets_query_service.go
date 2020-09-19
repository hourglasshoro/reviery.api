package twitter

import (
	"encoding/json"
	"fmt"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/query_service"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/query_service/read_model"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const MaxTweetsResultNumber = 50

type FetchTweetsQueryService struct {
}

func NewFetchTweetsQueryService() *FetchTweetsQueryService {
	qs := new(FetchTweetsQueryService)
	return qs
}

type TweetsResponse struct {
	Data []struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
	Meta struct {
		NewestId    string `json:"newest_id"`
		OldestId    string `json:"oldest_id"`
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	} `json:"meta"`
}

func (qs *FetchTweetsQueryService) FetchTweets(keyword value_object.Keyword, token value_object.AccessToken) (readModel read_model.FetchTweetsReadModel, err error) {
	u := fmt.Sprintf("https://api.twitter.com/2/tweets/search/recent?max_results=%d&query=%s", MaxTweetsResultNumber, url.QueryEscape(keyword.Keyword))
	log.Print(u)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		log.Print(res.Status)
		err = query_service.CannotFetchTweetsException
		return
	}
	defer res.Body.Close()

	var tweetsRes TweetsResponse
	err = json.NewDecoder(res.Body).Decode(&tweetsRes)
	if err != nil {
		return
	}
	var tweets []read_model.TweetReadModel
	for _, v := range tweetsRes.Data {
		id, e := strconv.Atoi(v.Id)
		if e != nil {
			err = e
			return
		}
		tweets = append(tweets, read_model.TweetReadModel{
			Id:   uint64(id),
			Text: v.Text,
		})
	}
	readModel = read_model.FetchTweetsReadModel{Tweets: tweets}
	return
}
