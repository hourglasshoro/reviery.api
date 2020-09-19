package cotoha

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hourglasshoro/reviery.api/src/app/domain/entitiy"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/query_service/read_model"
	"log"
	"net/http"
	"sync"
)

type FetchSentimentQueryService struct {
}

func NewFetchSentimentQueryService() *FetchSentimentQueryService {
	qs := new(FetchSentimentQueryService)
	return qs
}

type SentimentResponse struct {
	Result struct {
		Sentiment       string  `json:"sentiment"`
		Score           float64 `json:"score"`
		EmotionalPhrase []struct {
			Form    string `json:"form"`
			Emotion string `json:"emotion"`
		} `json:"emotional_phrase"`
	} `json:"result"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (qs *FetchSentimentQueryService) FetchSentiment(opinions []entitiy.Opinion, token value_object.AccessToken) (readModel read_model.FetchSentimentsReadModel, err error) {
	url := "https://api.ce-cotoha.com/api/dev/nlp/v1/sentiment"

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	var results []read_model.SentimentReadModel

	for _, opi := range opinions {
		wg.Add(1)
		go func(opinion entitiy.Opinion) {

			defer wg.Done()
			body := map[string]string{
				"sentence": opinion.Content,
			}
			b, err := json.Marshal(body)
			if err != nil {
				return
			}
			req, err := http.NewRequest("POST", url, bytes.NewReader(b))
			if err != nil {
				return
			}
			req.Header.Set("Content-Type", "application/json;charset=UTF-8")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))
			client := new(http.Client)
			res, err := client.Do(req)
			if err != nil {
				return
			}
			if res.StatusCode != http.StatusOK {
				err = errors.New(res.Status)
				log.Print(err)
				return
			}
			defer res.Body.Close()

			var senti SentimentResponse
			err = json.NewDecoder(res.Body).Decode(&senti)
			if err != nil {
				return
			}

			var sentiType value_object.SentimentType
			switch senti.Result.Sentiment {
			case "Positive":
				sentiType = value_object.SentimentTypes.Positive
			case "Negative":
				sentiType = value_object.SentimentTypes.Negative
			case "Neutral":
				sentiType = value_object.SentimentTypes.Neutral
			}

			mutex.Lock()

			results = append(results, read_model.SentimentReadModel{
				Id:    opinion.Id,
				Text:  opinion.Content,
				Score: float32(senti.Result.Score),
				Type:  sentiType,
			})

			mutex.Unlock()
		}(opi)
	}
	wg.Wait()

	readModel = read_model.FetchSentimentsReadModel{List: results}
	return
}
