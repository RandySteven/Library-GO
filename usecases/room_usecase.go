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
	"mime/multipart"
	"sync"
	"time"
)

type roomUsecase struct {
	roomRepo      repositories_interfaces.RoomRepository
	roomPhotoRepo repositories_interfaces.RoomPhotoRepository
	awsClient     aws_client.AWS
}

func (r *roomUsecase) CreateRoom(ctx context.Context, request *requests.CreateRoomRequest, fileHeader *multipart.FileHeader) (result *responses.CreateRoomResponse, customErr *apperror.CustomError) {
	imagePath, err := r.awsClient.UploadImageFile(request.Thumbnail, "rooms/", fileHeader, 1020, 800)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to upload book image", err)
	}

	room, err := r.roomRepo.Save(ctx, &models.Room{
		Name:      request.Name,
		Thumbnail: *imagePath,
	})
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to create room`, err)
	}

	result = &responses.CreateRoomResponse{
		ID:        utils.HashID(room.ID),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	return result, nil
}

func (r *roomUsecase) GetAllRooms(ctx context.Context) (results []*responses.ListRoomResponse, customErr *apperror.CustomError) {
	rooms, err := r.roomRepo.FindAll(ctx, 0, 0)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get all rooms`, err)
	}

	for _, room := range rooms {
		results = append(results, &responses.ListRoomResponse{
			ID:          room.ID,
			RoomName:    room.Name,
			Thumbnail:   room.Thumbnail,
			IsAvailable: room.IsAvailable,
			CreatedAt:   room.CreatedAt,
			UpdatedAt:   room.UpdatedAt,
			DeletedAt:   room.DeletedAt,
		})
	}

	return results, nil
}

func (r *roomUsecase) GetRoomByID(ctx context.Context, id uint64) (result *responses.RoomDetailResponse, customErr *apperror.CustomError) {
	room, err := r.roomRepo.FindByID(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get room`, err)
	}

	result = &responses.RoomDetailResponse{
		ID:        room.ID,
		RoomName:  room.Name,
		Thumbnail: room.Thumbnail,
	}

	return result, nil
}

func (r *roomUsecase) UploadRoomPhoto(ctx context.Context, request *requests.UploadRoomPhoto, fileHeader *multipart.FileHeader) (result *responses.UploadRoomPhotoResponse, customErr *apperror.CustomError) {
	var (
		wg           sync.WaitGroup
		imagePathsCh = make(chan []string, 0)
		customErrCh  = make(chan *apperror.CustomError)
	)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg.Add(1)

	go func() {
		defer wg.Done()
		imagePaths := []string{}
		for _, photo := range request.Photos {
			imagePath, err := r.awsClient.UploadImageFile(photo, "rooms/", fileHeader, 1020, 800)
			if err != nil {
				customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to upload image to s3`, err)
				return
			}
			imagePaths = append(imagePaths, *imagePath)
		}
		imagePathsCh <- imagePaths
	}()

	go func() {
		wg.Wait()
		close(customErrCh)
	}()

	select {
	case customErr = <-customErrCh:
		cancel()
		return nil, customErr
	case <-ctx.Done():
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "context cancelled", ctx.Err())
	default:
		for _, roomPhotoUrl := range <-imagePathsCh {
			roomPhotoUrl, err := r.roomPhotoRepo.Save(ctx, &models.RoomPhoto{
				RoomID: request.RoomID,
				Photo:  roomPhotoUrl,
			})
			if err != nil {
				return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to upload the photo`, err)
			}
			result.PhotoURLs = append(result.PhotoURLs, roomPhotoUrl.Photo)
		}
		result.RoomID = request.RoomID
		return result, nil
	}
}

var _ usecases_interfaces.RoomUsecase = &roomUsecase{}

func newRoomUsecase(
	roomRepo repositories_interfaces.RoomRepository,
	roomPhotoRepo repositories_interfaces.RoomPhotoRepository,
	awsClient aws_client.AWS) *roomUsecase {
	return &roomUsecase{
		roomRepo:      roomRepo,
		roomPhotoRepo: roomPhotoRepo,
		awsClient:     awsClient,
	}
}
