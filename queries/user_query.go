package queries

const (
	InsertUserQuery GoQuery = `
		INSERT INTO users (name, address, email, phone_number, password, dob)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	SelectUserByEmailQuery GoQuery = `
		SELECT 
		    id, name, address, email, phone_number, password, dob, profile_picture, created_at, updated_at, deleted_at, verified_at
		FROM users WHERE email = ?
	`

	SelectUserByPhoneNumberQuery GoQuery = `
		SELECT 
		    id, name, address, email, phone_number, password, dob, profile_picture, created_at, updated_at, deleted_at, verified_at
		FROM users WHERE phone_number = ?
	`

	SelectUserByIDQuery GoQuery = `
		SELECT 
		    id, name, address, email, phone_number, password, dob, profile_picture, created_at, updated_at, deleted_at, verified_at 
		FROM users WHERE id = ?
	`
)
