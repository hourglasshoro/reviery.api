package query_service

import (
	"errors"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
)

var NoCotohaAccessTokenExistException = errors.New("no Cotoha access token exists")

type GetCotohaAccessTokenQueryService interface {
	GetCotohaToken() (token value_object.AccessToken, err error)
}
