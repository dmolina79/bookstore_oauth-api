package access_token

import (
	"github.com/dmolina79/bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	defaultExpirationTime      = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	// used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`
	// used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (ar *AccessTokenRequest) Validate() *errors.RestErr {
	if ar.GrantType != grantTypePassword || ar.GrantType != grantTypeClientCredentials {
		return errors.NewBadRequest("invalid grant type")
	}
	// TODO: validate parameters for each grant type
	return nil
}

// TODO: generate code for JWT
func GetNewAccessToken(userId string) AccessToken {
	return AccessToken{
		AccessToken: userId,
		UserId:      0,
		ClientId:    0,
		Expires:     time.Now().UTC().Add(defaultExpirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequest("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequest("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequest("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequest("invalid expiration time")
	}

	return nil
}
