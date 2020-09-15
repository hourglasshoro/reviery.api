package twitter

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type AuthorizeQueryService struct {
}

func NewAuthorizeQueryService() *AuthorizeQueryService {
	qs := new(AuthorizeQueryService)
	return qs
}

type AuthResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

func (qs *AuthorizeQueryService) Auth() (token string, err error) {
	url := "https://api.twitter.com/oauth2/token?grant_type=client_credentials"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return
	}
	credential := []byte(fmt.Sprintf("%s:%s", os.Getenv("TWITTER_API_KEY"), os.Getenv("TWITTER_API_SECRET")))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(credential)))
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return
	}
	if res.Status != strconv.Itoa(http.StatusOK) {
		err = errors.New("not ok")
	}
	defer res.Body.Close()

	var auth AuthResponse
	err = json.NewDecoder(res.Body).Decode(&auth)
	if err != nil {
		return
	}

	return auth.AccessToken, nil
}
