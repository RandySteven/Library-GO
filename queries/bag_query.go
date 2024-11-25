package queries

const (
	InsertBagQuery GoQuery = `
		INSERT INTO bags (user_id, book_id)
		VALUES 
		    (?, ?)
	`

	SelectBagByUserQuery GoQuery = `
		SELECT b.id, b.user_id, b.book_id, 
		       u.id, u.name, u.address, u.email, u.phone_number, u.password, u.dob, u.profile_picture, 
		       u.created_at, u.updated_at, u.deleted_at, u.verified_at,
		       bo.id, bo.title, bo.description, bo.image, bo.status, bo.created_at, bo.updated_at, bo.deleted_at
			FROM bags AS b
				INNER JOIN users AS u
			ON
				b.user_id = u.id
				INNER JOIN books AS bo
			ON 
				b.book_id = bo.id
		WHERE b.user_id = ?
	`

	SelectExistBookAlreadyInBag GoQuery = `
		SELECT EXISTS(SELECT * FROM bags WHERE book_id = ? AND user_id = ?)
	`

	DeleteUserBagQuery GoQuery = `
		DELETE FROM bags WHERE user_id = ?
	`
)
