package repository

import (
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	"time"
)

type CotohaAccessTokenRepository interface {
	SetCotohaToken(token value_object.AccessToken, ttl time.Duration) (err error)
}
