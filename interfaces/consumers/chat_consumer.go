package consumers_interfaces

import "context"

type ChatConsumer interface {
	SendChatNotification(ctx context.Context)
	ReceiveChatNotification(ctx context.Context)
}
