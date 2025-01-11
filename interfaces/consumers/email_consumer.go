package consumers_interfaces

import "context"

type EmailConsumers interface {
	SendEmailNotification(ctx context.Context)
}
