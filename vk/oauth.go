package vk

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

type OAuthService struct {
	Client *Client
	Logger *log.Logger

	ClientID     int
	ClientSecret string
}

type AccessToken struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   int     `json:"expires_in"`
	UserID      int64   `json:"user_id"`
	Email       *string `json:"email,omitempty"`

	Error            *string `json:"error,omitempty"`
	ErrorDescription *string `json:"error_description,omitempty"`
}

// Get OAuth user access token.
// See https://vk.com/dev/authcode_flow_user
func (s *OAuthService) AccessToken(redirectURI string, code string) (*AccessToken, *Response, error) {
	u := fmt.Sprintf(
		"access_token?client_id=%d&client_secret=%s&redirect_uri=%s&code=%s",
		s.ClientID, s.ClientSecret, redirectURI, code,
	)

	fmt.Printf("%s", u)

	req, err := s.Client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var aToken = new(AccessToken)
	resp, err := s.Client.Do(req, aToken)
	if err != nil {
		return nil, resp, err
	}

	if aToken.Error != nil {
		err := errors.New(*aToken.Error)
		return nil, resp, err
	}

	return aToken, resp, err
}
