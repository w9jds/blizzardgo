package blizzard

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// AccessToken is an instance of an Access Token for the Blizzard Api
type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint64 `json:"expires_in"`
	Scope       string `json:"scope,omitempty"`
	ExpiresAt   time.Time
}

// UserInfo contains information about the owner of the access token
type UserInfo struct {
	Sub       string `json:"sub"`
	ID        uint32 `json:"id"`
	Battletag string `json:"battletag"`
}

func updateTokenTimes(token *AccessToken) *AccessToken {
	token.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn - 60))
	return token
}

// GetUserAccessToken gets the access token for the authenticated user
func (api API) GetUserAccessToken(code string) (*AccessToken, error) {
	data := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {code},
		"redirect_uri": {api.redirectURI},
	}

	request, err := http.NewRequest("POST", api.authURI+"/oauth/token", strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(api.clientID, api.clientSecret)
	if err != nil {
		return nil, err
	}

	var token *AccessToken
	err = api.do(request, &token)
	if err != nil {
		return nil, err
	}

	return updateTokenTimes(token), nil
}

// GetClientAccessToken gets a token for the current client application
func (api API) GetClientAccessToken() (*AccessToken, error) {
	request, err := http.NewRequest("GET", api.authURI+"/oauth/token?grant_type=client_credentials", nil)
	request.SetBasicAuth(api.clientID, api.clientSecret)
	if err != nil {
		return nil, err
	}

	var token *AccessToken
	err = api.do(request, &token)
	if err != nil {
		return nil, err
	}

	return updateTokenTimes(token), nil
}

// GetUserInfo returns the information about the user associated to the passed in token
func (api API) GetUserInfo(token AccessToken) (*UserInfo, error) {
	request, err := http.NewRequest("GET", api.authURI+"/oauth/userinfo", nil)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	if err != nil {
		return nil, err
	}

	var info *UserInfo
	err = api.do(request, &info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

// GetUserAuthorizationURI get the redirect url used to get the user to SSO into their account
func (api API) GetUserAuthorizationURI(scopes string) string {
	return fmt.Sprintf(
		"%s/oauth/authorize?client_id=%s&scope=%s&redirect_uri=%s&response_type=%s",
		api.authURI,
		api.clientID,
		scopes,
		api.redirectURI,
		"code",
	)
}
