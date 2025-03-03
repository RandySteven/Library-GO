package repositories

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type commentRepository struct {
	dbx repositories_interfaces.DB
}

func (c *commentRepository) DeleteByID(ctx context.Context, id uint64) error {
	return utils.Delete[models.Comment](ctx, c.dbx(ctx), `comments`, id)
}

func (c *commentRepository) Save(ctx context.Context, entity *models.Comment) (result *models.Comment, err error) {
	id, err := utils.Save[models.Comment](ctx, c.dbx(ctx), queries.InsertCommentQuery, &entity.UserID, &entity.BookID, &entity.ParentID, &entity.Comment)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (c *commentRepository) FindByID(ctx context.Context, id uint64) (result *models.Comment, err error) {
	result = &models.Comment{}
	err = utils.FindByID[models.Comment](ctx, c.dbx(ctx), queries.SelectCommentByIDQuery, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *commentRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *commentRepository) FindCommentsByBookID(ctx context.Context, bookID uint64) (result []*models.Comment, err error) {
	rows, err := c.dbx(ctx).QueryContext(ctx, queries.SelectBookCommentsQuery.ToString(), bookID)
	for rows.Next() {
		comment := &models.Comment{}
		err = rows.Scan(
			&comment.ID, &comment.UserID, &comment.BookID, &comment.ParentID, &comment.Comment,
			&comment.CreatedAt, &comment.UpdatedAt, &comment.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, comment)
	}
	return result, nil
}

var _ repositories_interfaces.CommentRepository = &commentRepository{}

func newCommentRepository(dbx repositories_interfaces.DB) *commentRepository {
	return &commentRepository{
		dbx: dbx,
	}
}
