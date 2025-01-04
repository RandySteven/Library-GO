package oauth2_client

import (
	"github.com/RandySteven/Library-GO/pkg/configs"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Oauth2Client struct {
	config oauth2.Config
}

func NewOauth2Client(config *configs.Config) (client *Oauth2Client) {
	oauth2Config := config.Config.Oauth2

	configOauth := oauth2.Config{
		RedirectURL:  oauth2Config.RedirectURL,
		ClientID:     oauth2Config.ClientID,
		ClientSecret: oauth2Config.ClientSecret,
		Scopes:       oauth2Config.Scopes,
		Endpoint:     google.Endpoint,
	}

	return &Oauth2Client{
		config: configOauth,
	}
}
