package main

import (
	"github.com/hourglasshoro/reviery.api/src/app/domain/domain_service"
	"github.com/hourglasshoro/reviery.api/src/app/infrastructure/redis"
	"github.com/hourglasshoro/reviery.api/src/app/infrastructure/web/cotoha"
	"github.com/hourglasshoro/reviery.api/src/app/infrastructure/web/twitter"
	"github.com/hourglasshoro/reviery.api/src/app/presentation/controller"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/interactor"
)

type Controllers struct {
	Common controller.CommonController
}

func NewControllers() (ctrls *Controllers, err error) {
	ctrls = new(Controllers)
	redisInst, err := redis.NewRedis()
	if err != nil {
		return
	}

	twitterRepo := redis.NewTwitterAccessTokenRepository(redisInst)
	getTwitterToken := redis.NewGetTwitterAccessTokenQueryService(redisInst)
	cotohaRepo := redis.NewCotohaAccessTokenRepository(redisInst)
	getCotohaToken := redis.NewGetCotohaAccessTokenQueryService(redisInst)
	twitterAuth := twitter.NewAuthorizeTwitterQueryService()
	cotohaAuth := cotoha.NewAuthorizeCotohaQueryService()
	fetchTweets := twitter.NewFetchTweetsQueryService()
	fetchSentiments := cotoha.NewFetchSentimentQueryService()

	calcScore := *domain_service.NewCalcSentimentScoreDomainService()

	searchIntr := *interactor.NewSearchInteractor(
		twitterRepo,
		getTwitterToken,
		cotohaRepo,
		getCotohaToken,
		twitterAuth,
		cotohaAuth,
		fetchTweets,
		fetchSentiments,
		calcScore,
	)

	commonCtrl := controller.NewCommonController(searchIntr)
	ctrls.Common = *commonCtrl
	return ctrls, nil
}
