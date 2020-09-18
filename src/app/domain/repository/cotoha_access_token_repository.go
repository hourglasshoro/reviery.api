package repository

import "github.com/hourglasshoro/reviery.api/src/app/domain/value_object"

type CotohaAccessTokenRepository interface {
	SetCotohaToken(token value_object.AccessToken) (err error)
}
