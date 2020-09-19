package controller

import (
	"context"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	pb "github.com/hourglasshoro/reviery.api/src/app/presentation/grpc/common"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/interactor"
	"sync"
)

type CommonController struct {
	search interactor.SearchInteractor
}

func NewCommonController(search interactor.SearchInteractor) *CommonController {
	ctrl := new(CommonController)
	ctrl.search = search
	return ctrl
}

func (ctrl *CommonController) Search(ctx context.Context, in *pb.SearchMessage) (response *pb.SearchResponse, err error) {
	keyword, err := value_object.NewKeyword(in.Keyword)
	if err != nil {
		return
	}
	output, err := ctrl.search.Invoke(*keyword)
	if err != nil {
		return
	}

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	var opinions []*pb.Opinion

	for _, v := range output.Opinions {
		wg.Add(1)
		go func(item value_object.SentimentResult) {
			defer wg.Done()
			var typ pb.SentimentType
			switch item.Type {
			case value_object.SentimentTypes.Positive:
				typ = pb.SentimentType_Positive
			case value_object.SentimentTypes.Negative:
				typ = pb.SentimentType_Negative
			case value_object.SentimentTypes.Neutral:
				typ = pb.SentimentType_Neutral
			}

			mutex.Lock()
			opinions = append(opinions, &pb.Opinion{
				Id:    item.Id,
				Text:  item.Text,
				Score: item.Score,
				Type:  typ,
			})
			mutex.Unlock()
		}(v)
	}
	wg.Wait()

	return &pb.SearchResponse{
		Opinions:   opinions,
		TotalScore: output.TotalScore,
	}, nil
}
