package access_token

import (
	"github.com/dmolina79/bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	defaultExpirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		AccessToken: "",
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
