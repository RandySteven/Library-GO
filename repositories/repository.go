package repositories

import (
	"database/sql"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
)

type Repositories struct {
	AuthorRepo       repositories_interfaces.AuthorRepository
	AuthorBookRepo   repositories_interfaces.AuthorBookRepository
	BagRepo          repositories_interfaces.BagRepository
	BookRepo         repositories_interfaces.BookRepository
	BookGenreRepo    repositories_interfaces.BookGenreRepository
	BorrowRepo       repositories_interfaces.BorrowRepository
	BorrowDetailRepo repositories_interfaces.BorrowDetailRepository
	GenreRepo        repositories_interfaces.GenreRepository
	UserRepo         repositories_interfaces.UserRepository
	RatingRepo       repositories_interfaces.RatingRepository
	CommentRepo      repositories_interfaces.CommentRepository
	EventRepo        repositories_interfaces.EventRepository
	EventUserRepo    repositories_interfaces.EventUserRepository
	ChatRepo         repositories_interfaces.ChatRepository
	RoomChatRepo     repositories_interfaces.RoomChatRepository
	RoomChatUserRepo repositories_interfaces.RoomChatUserRepository
	RoleUserRepo     repositories_interfaces.RoleUserRepository
	RoomRepo         repositories_interfaces.RoomRepository
	RoomPhotoRepo    repositories_interfaces.RoomPhotoRepository
	Transaction      repositories_interfaces.Transaction
}

func NewRepositories(db *sql.DB) *Repositories {
	tr, dbx := newTransaction(db)
	return &Repositories{
		AuthorRepo:       newAuthorRepository(dbx),
		AuthorBookRepo:   newAuthorBookRepository(dbx),
		BagRepo:          newBagRepository(dbx),
		BookRepo:         newBookRepository(dbx),
		BookGenreRepo:    newBookGenreRepository(dbx),
		BorrowRepo:       newBorrowRepository(dbx),
		BorrowDetailRepo: newBorrowDetailRepository(dbx),
		GenreRepo:        newGenreRepository(dbx),
		UserRepo:         newUserRepository(dbx),
		RatingRepo:       newRatingRepository(dbx),
		CommentRepo:      newCommentRepository(dbx),
		EventRepo:        newEventRepository(dbx),
		EventUserRepo:    newEventUserRepository(dbx),
		ChatRepo:         newChatRepository(dbx),
		RoomChatRepo:     newRoomChatRepository(dbx),
		RoomChatUserRepo: newRoomChatUserRepository(dbx),
		RoleUserRepo:     newRoleUserRepository(dbx),
		RoomRepo:         newRoomRepository(dbx),
		RoomPhotoRepo:    newRoomPhotoRepository(dbx),
		Transaction:      tr,
	}
}
