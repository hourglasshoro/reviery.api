package query_service

import (
	"errors"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
)

var CannotGetCotohaAccessTokenException = errors.New("cannot get cotoha access token")

type AuthorizeCotohaQueryService interface {
	AuthCotoha() (token value_object.AccessToken, err error)
}
