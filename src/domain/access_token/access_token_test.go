package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, defaultExpirationTime, 24)
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	assert.NotNil(t, at, "new access token should not be nil")
	assert.False(t, at.IsExpired(), "new access token should not expired")
	assert.Empty(t, at.AccessToken, "new access token should not have defined id")
	assert.Empty(t, at.UserId, "new access token should not have a user id")
}

func TestAccessToken_IsExpired_AsDefault(t *testing.T) {
	at := AccessToken{}

	assert.True(t, at.IsExpired(), "empty access token should be expired by default")
}

func TestAccessToken_IsExpired_OnValidExpirationDate(t *testing.T) {
	at := AccessToken{}
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()

	assert.False(t, at.IsExpired(), "should not expire with 3 hours expiration time")
}
