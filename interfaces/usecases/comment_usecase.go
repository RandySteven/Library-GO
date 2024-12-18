package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type CommentUsecase interface {
	AddComment(ctx context.Context, request *requests.AddCommentRequest) (result *responses.CommentResponse, customErr *apperror.CustomError)
	ReplyComment(ctx context.Context, request *requests.ReplyCommentRequest) (result *responses.ReplyCommentResponse, customErr *apperror.CustomError)
	GetCommentFromBook(ctx context.Context, request *requests.GetCommentRequest) (result []*responses.ListBookCommentsResponse, customErr *apperror.CustomError)
	DeleteComment(ctx context.Context, id uint64) (customErr *apperror.CustomError)
}
