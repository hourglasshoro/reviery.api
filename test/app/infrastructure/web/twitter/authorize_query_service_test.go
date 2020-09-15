package test_twitter

import (
	"github.com/hourglasshoro/reviery.api/src/app/infrastructure/web/twitter"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAuthorizeQueryService(t *testing.T) {
	err := godotenv.Load("../../../../../.env")
	if err != nil {
		t.Error(err)
	}

	qs := twitter.NewAuthorizeQueryService()
	token, err := qs.Auth()
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	log.Print(token)
}
