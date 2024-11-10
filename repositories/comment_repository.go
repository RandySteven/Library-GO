package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type commentRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (c *commentRepository) Save(ctx context.Context, entity *models.Comment) (result *models.Comment, err error) {
	id, err := utils.Save[models.Comment](ctx, c.InitTrigger(), queries.InsertCommentQuery, &entity.UserID, &entity.BookID, &entity.ParentID, &entity.Comment)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (c *commentRepository) FindByID(ctx context.Context, id uint64) (result *models.Comment, err error) {
	result = &models.Comment{}
	err = utils.FindByID[models.Comment](ctx, c.InitTrigger(), queries.SelectCommentByIDQuery, id, result)
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
	rows, err := c.InitTrigger().QueryContext(ctx, queries.SelectBookCommentsQuery.ToString(), bookID)
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

func (c *commentRepository) InitTrigger() repositories_interfaces.Trigger {
	var trigger repositories_interfaces.Trigger = c.db
	if c.tx != nil {
		trigger = c.tx
	}
	return trigger
}

func (c *commentRepository) BeginTx(ctx context.Context) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	c.tx = tx
	return nil
}

func (c *commentRepository) CommitTx(ctx context.Context) error {
	return c.tx.Commit()
}

func (c *commentRepository) RollbackTx(ctx context.Context) error {
	return c.tx.Rollback()
}

func (c *commentRepository) SetTx(tx *sql.Tx) {
	c.tx = tx
}

func (c *commentRepository) GetTx(ctx context.Context) *sql.Tx {
	return c.tx
}

var _ repositories_interfaces.CommentRepository = &commentRepository{}

func newCommentRepository(db *sql.DB) *commentRepository {
	return &commentRepository{
		db: db,
	}
}
