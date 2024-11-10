package caches_interfaces

import "context"

type ChatCache interface {
	SendChat(ctx context.Context)
	GetChatID(ctx context.Context)
}
