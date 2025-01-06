package queries

const (
	InsertRoleUserQuery GoQuery = `
		INSERT INTO role_users (role_id, user_id)
		VALUES 
		    (?, ?)
	`

	SelectRoleUserByUserIDQuery GoQuery = `
		SELECT id, role_id, user_id, created_at, updated_at, deleted_at 
			FROM role_users
		WHERE 
		    user_id = ?
	`
)
