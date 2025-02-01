package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
	"github.com/RandySteven/Library-GO/utils"
	"io"
	"mime/multipart"
	"time"
)

type eventUsecase struct {
	awsClient     aws_client.AWS
	eventRepo     repositories_interfaces.EventRepository
	eventUserRepo repositories_interfaces.EventUserRepository
	userRepo      repositories_interfaces.UserRepository
}

func (e *eventUsecase) imageUploader(ctx context.Context, thumbnail io.Reader, fileHeader *multipart.FileHeader) (*string, *apperror.CustomError) {
	imagePath, err := e.awsClient.UploadImageFile(ctx, thumbnail, "events/", fileHeader, 1544, 794)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to upload book image", err)
	}
	return imagePath, nil
}

func (e *eventUsecase) CreateEvent(ctx context.Context, request *requests.CreateEventRequest, fileHeader *multipart.FileHeader) (result *responses.EventCreateResponse, customErr *apperror.CustomError) {
	event := &models.Event{
		Title:             request.Title,
		Price:             request.Price,
		Description:       request.Description,
		ParticipantNumber: request.ParticipantNumber,
	}

	imagePath, customErr := e.imageUploader(ctx, request.Thumbnail, fileHeader)
	if customErr != nil {
		return nil, customErr
	}
	event.Thumbnail = *imagePath

	event, err := e.eventRepo.Save(ctx, event)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to insert event`, err)
	}

	return &responses.EventCreateResponse{
		ID:        utils.HashID(event.ID),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		DeletedAt: nil,
	}, nil
}

func (e *eventUsecase) GetAllEvents(ctx context.Context) (result []*responses.ListEventResponse, customErr *apperror.CustomError) {
	events, err := e.eventRepo.FindAll(ctx, 0, 0)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get all events`, err)
	}

	for _, event := range events {
		result = append(result, &responses.ListEventResponse{
			ID:        event.ID,
			Title:     event.Title,
			Thumbnail: event.Thumbnail,
			Price:     event.Price,
			Slot:      event.OccupiedParticipantNumber - event.ParticipantNumber,
			Date:      event.Date,
			StartTime: event.StartTime,
			EndTime:   event.EndTime,
			CreatedAt: event.CreatedAt,
			DeletedAt: event.DeletedAt,
		})
	}

	return result, nil
}

func (e *eventUsecase) GetEvent(ctx context.Context, id uint64) (result *responses.EventDetailResponse, customErr *apperror.CustomError) {
	event, err := e.eventRepo.FindByID(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get event`, err)
	}
	result = &responses.EventDetailResponse{
		ID:          event.ID,
		Title:       event.Title,
		Thumbnail:   event.Thumbnail,
		Description: event.Description,
		Price:       event.Price,
		Slot:        event.OccupiedParticipantNumber - event.ParticipantNumber,
		Date:        event.Date,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
		DeletedAt:   event.DeletedAt,
	}
	return result, nil
}

var _ usecases_interfaces.EventUsecase = &eventUsecase{}

func newEventUsecase(
	awsClient aws_client.AWS,
	eventRepo repositories_interfaces.EventRepository,
	eventUserRepo repositories_interfaces.EventUserRepository,
	userRepo repositories_interfaces.UserRepository,
) *eventUsecase {
	return &eventUsecase{
		awsClient:     awsClient,
		eventRepo:     eventRepo,
		eventUserRepo: eventUserRepo,
		userRepo:      userRepo,
	}
}
