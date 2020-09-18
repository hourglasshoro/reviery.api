package cotoha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hourglasshoro/reviery.api/src/app/domain/entitiy"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/query_service/read_model"
	"net/http"
	"sync"
)

type FetchSentimentQueryService struct {
}

type SentimentResponse struct {
	Result struct {
		Sentiment       string  `json:"sentiment"`
		Score           float32 `json:"score"`
		EmotionalPhrase []struct {
			Form    string `json:"form"`
			Emotion string `json:"emotion"`
		} `json:"emotional_phrase"`
	} `json:"result"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (qs *FetchSentimentQueryService) FetchSentiment(opinions []entitiy.Opinion, token value_object.AccessToken) (readModel read_model.FetchSentimentsReadModel, err error) {
	url := "https://api.ce-cotoha.com/v1/nlp/v1/sentiment"

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
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			client := new(http.Client)
			res, err := client.Do(req)
			if err != nil {
				return
			}
			defer res.Body.Close()

			var senti SentimentResponse
			err = json.NewDecoder(res.Body).Decode(&senti)
			if err != nil {
				return
			}
			var sentiType read_model.SentimentType
			switch senti.Result.Sentiment {
			case "Positive":
				sentiType = read_model.SentimentTypes.Positive
			case "Negative":
				sentiType = read_model.SentimentTypes.Negative
			case "Neutral":
				sentiType = read_model.SentimentTypes.Neutral
			}
			mutex.Lock()
			results = append(results, read_model.SentimentReadModel{
				Id:    opinion.Id,
				Text:  opinion.Content,
				Score: senti.Result.Score,
				Type:  sentiType,
			})
			mutex.Unlock()
		}(opi)
	}

	readModel = read_model.FetchSentimentsReadModel{List: results}
	return
}
