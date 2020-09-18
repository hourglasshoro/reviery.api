package repository

import "github.com/hourglasshoro/reviery.api/src/app/domain/value_object"

type TwitterAccessTokenRepository interface {
	SetTwitterToken(token value_object.AccessToken) (err error)
}
