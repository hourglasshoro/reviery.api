package cotoha

import (
	"bytes"
	"encoding/json"
	"github.com/hourglasshoro/reviery.api/src/app/domain/value_object"
	"github.com/hourglasshoro/reviery.api/src/app/usecase/query_service"
	"net/http"
	"os"
)

type AuthorizeQueryService struct {
}

func NewAuthorizeCotohaQueryService() *AuthorizeQueryService {
	qs := new(AuthorizeQueryService)
	return qs
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"score"`
	IssuedAt    string `json:"issued_at"`
}

func (qs *AuthorizeQueryService) AuthCotoha() (token value_object.AccessToken, err error) {
	url := "https://api.ce-cotoha.com/v1/oauth/accesstokens"
	body := map[string]string{
		"grantType":    "client_credentials",
		"clientId":     os.Getenv("COTOHA_API_KEY"),
		"clientSecret": os.Getenv("COTOHA_API_SECRET"),
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
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		err = query_service.CannotGetCotohaAccessTokenException
	}
	defer res.Body.Close()

	var auth AuthResponse
	err = json.NewDecoder(res.Body).Decode(&auth)
	if err != nil {
		return
	}
	tok, err := value_object.NewAccessToken(auth.AccessToken)
	if err != nil {
		return
	}
	token = *tok
	return
}
