package queries

const (
	InsertBagQuery GoQuery = `
		INSERT INTO bags (user_id, book_id)
		VALUES 
		    (?, ?)
	`

	SelectBagByUserQuery GoQuery = `
		SELECT id, user_id, book_id
		FROM bags
		WHERE user_id = ?
	`

	SelectExistBookAlreadyInBag GoQuery = `
		SELECT EXISTS(SELECT * FROM bags WHERE book_id = ? AND user_id = ?)
	`

	DeleteUserBagQuery GoQuery = `
		DELETE FROM bags WHERE user_id = ?
	`
)
