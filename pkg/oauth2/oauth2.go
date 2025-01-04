package oauth2_client

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Oauth2Client struct {
	config oauth2.Config
}

func NewOauth2Client() (client *Oauth2Client) {
	config := oauth2.Config{
		RedirectURL:  "",
		ClientID:     "",
		ClientSecret: "",
		Scopes:       nil,
		Endpoint:     google.Endpoint,
	}

	return &Oauth2Client{
		config: config,
	}
}
