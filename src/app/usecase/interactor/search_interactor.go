package interactor

import (
	"github.com/hourglasshoro/reviery.api/src/app/domain/domain_service"
	"github.com/hourglasshoro/reviery.api/src/app/domain/entitiy"
	"github.com/hourglasshoro/reviery.api/src/app/domain/repository"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/query_service"
	"log"
	"time"
)

const AccessTokenTTL = time.Hour

type SearchOutput struct {
	TotalScore float32
	Opinions   []value_object.SentimentResult
}

type SearchInteractor struct {
	twitterRepo     repository.TwitterAccessTokenRepository
	getTwitterToken query_service.GetTwitterAccessTokenQueryService
	cotohaRepo      repository.CotohaAccessTokenRepository
	getCotohaToken  query_service.GetCotohaAccessTokenQueryService
	twitterAuth     query_service.AuthorizeTwitterQueryService
	cotohaAuth      query_service.AuthorizeCotohaQueryService
	fetchTweets     query_service.FetchTweetsQueryService
	fetchSentiments query_service.FetchSentimentQueryService
	calcScore       domain_service.CalcSentimentScoreDomainService
}

func NewSearchInteractor(
	twitterRepo repository.TwitterAccessTokenRepository,
	getTwitterToken query_service.GetTwitterAccessTokenQueryService,
	cotohaRepo repository.CotohaAccessTokenRepository,
	getCotohaToken query_service.GetCotohaAccessTokenQueryService,
	twitterAuth query_service.AuthorizeTwitterQueryService,
	cotohaAuth query_service.AuthorizeCotohaQueryService,
	fetchTweets query_service.FetchTweetsQueryService,
	fetchSentiments query_service.FetchSentimentQueryService,
	calcScore domain_service.CalcSentimentScoreDomainService,
) *SearchInteractor {
	intr := new(SearchInteractor)
	intr.twitterRepo = twitterRepo
	intr.getTwitterToken = getTwitterToken
	intr.cotohaRepo = cotohaRepo
	intr.getCotohaToken = getCotohaToken
	intr.twitterAuth = twitterAuth
	intr.cotohaAuth = cotohaAuth
	intr.fetchTweets = fetchTweets
	intr.fetchSentiments = fetchSentiments
	intr.calcScore = calcScore
	return intr
}

func (uc *SearchInteractor) Invoke(keyword value_object.Keyword) (output SearchOutput, err error) {

	// Get access token for Twitter
	twitterToken, err := uc.getTwitterToken.GetTwitterToken()
	log.Print("Get access token for Twitter")

	if err == query_service.NoTwitterAccessTokenExistException {
		newToken, e := uc.twitterAuth.AuthTwitter()
		log.Print("Authorize Twitter")

		if e != nil {
			err = e
			return
		}
		err = uc.twitterRepo.SetTwitterToken(newToken, AccessTokenTTL)
		log.Print("Set Twitter access token")

		twitterToken, err = uc.getTwitterToken.GetTwitterToken()
		log.Print("Get access token for Twitter")
	}
	if err != nil {
		return
	}

	// Get access token for Cotoha
	cotohaToken, err := uc.getCotohaToken.GetCotohaToken()
	log.Print("Get access token for Cotoha")

	if err == query_service.NoCotohaAccessTokenExistException {
		newToken, e := uc.cotohaAuth.AuthCotoha()
		log.Print("Authorize Cotoha")

		if e != nil {
			err = e
			return
		}
		err = uc.cotohaRepo.SetCotohaToken(newToken, AccessTokenTTL)
		log.Print("Set Cotoha access token")

		cotohaToken, err = uc.getCotohaToken.GetCotohaToken()
		log.Print("Get access token for Cotoha")
	}
	if err != nil {
		return
	}

	// Fetch tweets
	tweets, err := uc.fetchTweets.FetchTweets(keyword, twitterToken)
	log.Print("Search tweets")

	if err != nil {
		return
	}

	// Convert opinions
	var opinions []entitiy.Opinion
	for _, tweet := range tweets.Tweets {
		opinion, err := entitiy.NewOpinion(tweet.Id, tweet.Text)
		if err != nil {
			continue
		}
		opinions = append(opinions, *opinion)
	}
	// Analyze sentiment
	results, err := uc.fetchSentiments.FetchSentiment(opinions, cotohaToken)
	log.Print("Analyze opinions")

	if err != nil {
		return
	}

	// Calc Total Score
	totalScore, calcResults, err := uc.calcScore.CalcScore(results)
	log.Print("Calc total score")

	output = SearchOutput{
		totalScore, calcResults,
	}
	return
}
