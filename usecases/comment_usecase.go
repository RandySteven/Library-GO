package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
)

type commentUsecase struct {
	commentRepo repositories_interfaces.CommentRepository
	userRepo    repositories_interfaces.UserRepository
	bookRepo    repositories_interfaces.BookRepository
}

func (c *commentUsecase) AddComment(ctx context.Context, request *requests.AddCommentRequest) (result *responses.CommentResponse, customErr *apperror.CustomError) {
	userId := ctx.Value(enums.UserID).(uint64)
	comment := &models.Comment{
		UserID:   userId,
		BookID:   request.BookID,
		Comment:  request.Comment,
		ParentID: nil,
	}
	comment, err := c.commentRepo.Save(ctx, comment)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to save comment`, err)
	}
	result = &responses.CommentResponse{
		ID:      comment.ID,
		UserID:  comment.UserID,
		BookID:  comment.BookID,
		Comment: comment.Comment,
	}
	return result, nil
}

func (c *commentUsecase) ReplyComment(ctx context.Context, request *requests.ReplyCommentRequest) (result *responses.ReplyCommentResponse, customErr *apperror.CustomError) {
	comment := &models.Comment{
		UserID:   ctx.Value(enums.UserID).(uint64),
		BookID:   request.BookID,
		Comment:  request.Comment,
		ParentID: &request.ParentID,
	}
	if request.ReplyID != nil {
		comment.ReplyID = request.ReplyID
	}

	comment, err := c.commentRepo.Save(ctx, comment)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to reply comment`, err)
	}

	result = &responses.ReplyCommentResponse{
		ID:        comment.ID,
		UserID:    comment.UserID,
		BookID:    comment.BookID,
		CommentID: *comment.ParentID,
		Comment:   comment.Comment,
	}

	return result, nil
}

func (c *commentUsecase) GetCommentFromBook(ctx context.Context, request *requests.GetCommentRequest) (result []*responses.ListBookCommentsResponse, customErr *apperror.CustomError) {
	comments, err := c.commentRepo.FindCommentsByBookID(ctx, request.BookID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get comment`, err)
	}
	for _, comment := range comments {
		user, err := c.userRepo.FindByID(ctx, comment.UserID)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed get user`, err)
		}
		result = append(result, &responses.ListBookCommentsResponse{
			ID:       comment.ID,
			BookID:   comment.BookID,
			ParentID: comment.ParentID,
			Comment:  comment.Comment,
			User: struct {
				ID    uint64 `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{ID: user.ID, Name: user.Name, Email: user.Email},
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}
	return result, nil
}

func (c *commentUsecase) DeleteComment(ctx context.Context, id uint64) (customErr *apperror.CustomError) {
	err := c.commentRepo.DeleteByID(ctx, id)
	if err != nil {
		return apperror.NewCustomError(apperror.ErrInternalServer, `failed to delete comment`, err)
	}
	return nil
}

var _ usecases_interfaces.CommentUsecase = &commentUsecase{}

func newCommentUsecase(
	commentRepo repositories_interfaces.CommentRepository,
	userRepo repositories_interfaces.UserRepository,
	bookRepo repositories_interfaces.BookRepository) *commentUsecase {
	return &commentUsecase{
		commentRepo: commentRepo,
		userRepo:    userRepo,
		bookRepo:    bookRepo,
	}
}
