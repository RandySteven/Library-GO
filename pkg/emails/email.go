package emails_client

import (
	"crypto/tls"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	Dialer   *gomail.Dialer
	Message  *gomail.Message
	Metadata map[string]interface{}
}

func NewMailtrap(config *configs.Config) (*Mailer, error) {
	mailtrap := config.Config.Mailtrap
	host := mailtrap.Host
	port := mailtrap.Port
	username := mailtrap.Username
	password := mailtrap.Password

	dialer := gomail.NewDialer(host, port, username, password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &Mailer{
		Dialer:  dialer,
		Message: gomail.NewMessage(),
	}, nil
}
