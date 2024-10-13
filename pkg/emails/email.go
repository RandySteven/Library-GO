package emails_client

import "context"

type EmailsClient interface {
	SendEmail(ctx context.Context, to, from, subject, html string) error
}
