package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"mime/multipart"
)

type EventUsecase interface {
	CreateEvent(ctx context.Context, request *requests.CreateEventRequest, fileHeader *multipart.FileHeader) (result *responses.EventCreateResponse, customErr *apperror.CustomError)
	GetAllEvents(ctx context.Context) (result []*responses.ListEventResponse, customErr *apperror.CustomError)
	GetEvent(ctx context.Context, id uint64) (result *responses.EventDetailResponse, customErr *apperror.CustomError)
}
