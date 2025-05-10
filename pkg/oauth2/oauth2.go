package oauth2_client

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/enums"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

type (
	Oauth2 interface {
		LoginAuth(ctx context.Context) string
		CallbackAuth(ctx context.Context) (string, error)
	}

	Oauth2Client struct {
		config oauth2.Config
	}
)

func (o *Oauth2Client) LoginAuth(ctx context.Context) string {
	return o.config.AuthCodeURL(os.Getenv("RANDOM_STATE_OAUTH2"))
}

func (o *Oauth2Client) CallbackAuth(ctx context.Context) (string, error) {
	if ctx.Value(enums.RandomState).(string) != os.Getenv("RANDOM_STATE_OAUTH2") {
		return "", fmt.Errorf(`random state is not same`)
	}

	token, err := o.config.Exchange(ctx, "")
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func NewOauth2Client(config *configs.Config) (client *Oauth2Client, err error) {
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
	}, nil
}
