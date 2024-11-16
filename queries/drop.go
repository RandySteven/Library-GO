package queries

const (
	DropUserTable         DropQuery = "DROP TABLE IF EXISTS users"
	DropAuthorTable       DropQuery = "DROP TABLE IF EXISTS authors"
	DropGenreTable        DropQuery = "DROP TABLE IF EXISTS genres"
	DropBookTable         DropQuery = "DROP TABLE IF EXISTS books"
	DropBookGenreTable    DropQuery = "DROP TABLE IF EXISTS book_genres"
	DropAuthorBookTable   DropQuery = "DROP TABLE IF EXISTS author_books"
	DropBagTable          DropQuery = "DROP TABLE IF EXISTS bags"
	DropBorrowTable       DropQuery = "DROP TABLE IF EXISTS borrows"
	DropBorrowDetailTable DropQuery = "DROP TABLE IF EXISTS borrow_details"
	DropRatingTable       DropQuery = "DROP TABLE IF EXISTS ratings"
	DropCommentTable      DropQuery = "DROP TABLE IF EXISTS comments"
	DropEventTable        DropQuery = "DROP TABLE IF EXISTS events"
	DropEventUserTable    DropQuery = "DROP TABLE IF EXISTS event_users"
	DropRoomChatTable     DropQuery = "DROP TABLE IF EXISTS room_chats"
	DropRoomChatUserTable DropQuery = "DROP TABLE IF EXISTS room_chat_users"
	DropChatTable         DropQuery = "DROP TABLE IF EXISTS chats"
)
