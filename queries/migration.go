package queries

const (
	AuthorMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS authors (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    name VARCHAR(128) NOT NULL,
		    nationality VARCHAR(128) NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL 
		)
	`
	UserMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS users (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    name VARCHAR(128) NOT NULL,
		    address TEXT NOT NULL,
		    email VARCHAR(128) NOT NULL UNIQUE,
		    phone_number VARCHAR(13) NOT NULL UNIQUE,
		    password VARCHAR(244) NOT NULL,
		    dob DATE NOT NULL,
		    profile_picture VARCHAR(244) NOT NULL DEFAULT '',
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL ,
		    verified_at TIMESTAMP DEFAULT NULL 
		)
	`
	GenreMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS genres (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    genre VARCHAR(128) NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL 		    
		)
	`
	BookMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS books (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    title VARCHAR(128) NOT NULL,
		    description TEXT NOT NULL,
		    image VARCHAR(244) NOT NULL,
		    status INT NOT NULL DEFAULT 0,
		    pdf_file VARCHAR(244) NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL 
		)
	`
	AuthorBookMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS author_books (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    author_id BIGINT NOT NULL,
		    book_id BIGINT NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL ,
		    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
		    FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE
		)
	`
	BookGenreMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS book_genres (
				id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
				book_id BIGINT NOT NULL,
				genre_id BIGINT NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				deleted_at TIMESTAMP DEFAULT NULL ,
				FOREIGN KEY(book_id) REFERENCES books(id) ON DELETE CASCADE,
				FOREIGN KEY(genre_id) REFERENCES genres(id) ON DELETE CASCADE
			)
	`
	BagMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS bags (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    user_id BIGINT NOT NULL,
		    book_id BIGINT NOT NULL,
		    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
		)
	`
	BorrowMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS borrows (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    user_id BIGINT NOT NULL,
		    borrow_reference CHAR(24) NOT NULL UNIQUE,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL,
		    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)
	`
	BorrowDetailMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS borrow_details (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    borrow_id BIGINT NOT NULL,
		    book_id BIGINT NOT NULL,
		    borrowed_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    returned_date TIMESTAMP NOT NULL,
		    verified_return_date TIMESTAMP DEFAULT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL,
			FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
		    FOREIGN KEY (borrow_id) REFERENCES borrows(id) ON DELETE CASCADE
		)
	`
	StoryGeneratorMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS story_generators (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    prompt TEXT NOT NULL,
		    result TEXT NOT NULL,
		    image VARCHAR(244) NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL
		)
	`

	RatingMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS ratings (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    book_id BIGINT NOT NULL,
		    user_id BIGINT NOT NULL,
		    score FLOAT NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL,
		    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
		    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)
	`

	CommentMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS comments (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    user_id BIGINT NOT NULL,
		    book_id BIGINT NOT NULL,
		    parent_id BIGINT DEFAULT NULL,
		    reply_id BIGINT DEFAULT NULL,
		    comment TEXT NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL,
		    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
		    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		    FOREIGN KEY (parent_id) REFERENCES comments(id) ON DELETE CASCADE,
		    FOREIGN KEY (reply_id) REFERENCES comments(id) ON DELETE CASCADE
		)
	`

	EventMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS events (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    title VARCHAR(122) NOT NULL,
		    thumbnail VARCHAR(244) NOT NULL,
		    price INT DEFAULT NULL,
		    description TEXT NOT NULL,
		    participant_number INT NOT NULL,
		    occupied_participant_number INT NOT NULL,
		    date DATE NOT NULL,
		    start_time TIME NOT NULL,
		    end_time TIME NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL
		)
	`

	EventUserMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS event_users (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    user_id BIGINT NOT NULL,
		    event_id BIGINT NOT NULL,
		    payed BOOLEAN DEFAULT FALSE,
		    event_code VARCHAR(16) NOT NULL,
		    redeem_status INT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL,		    
		    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
		)
	`

	RoomChatMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS room_chats (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    room_name VARCHAR(64) NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL
		)
	`

	RoomChatUserMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS room_chat_users (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    room_chat_id BIGINT NOT NULL,
		    user_id BIGINT NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL,
		    FOREIGN KEY (room_chat_id) REFERENCES room_chats(id) ON DELETE CASCADE,
		    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)
	`

	ChatMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS chats (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    room_chat_id BIGINT NOT NULL,
		    user_id BIGINT NOT NULL,
		    chat TEXT NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL,
		    FOREIGN KEY (room_chat_id) REFERENCES room_chats(id) ON DELETE CASCADE,
		    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)
	`

	RoleMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS roles (
			id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
			role VARCHAR(16) NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL
		)
	`

	RoleUserMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS role_users (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    user_id BIGINT NOT NULL,
		    role_id BIGINT NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL,
		    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
            FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
		)
	`

	UserGenreMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS user_genres (
			id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
			user_id BIGINT NOT NULL,
			genre_id BIGINT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		    FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
		)
	`

	RoomMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS rooms (
		    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		    name VARCHAR(36) NOT NULL UNIQUE,
		    thumbnail VARCHAR(244) NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL
		)
	`

	RoomPhotoMigration MigrationQuery = `
		CREATE TABLE IF NOT EXISTS room_photos (
			id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
			photo VARCHAR(244) NOT NULL,	    
		    room_id BIGINT NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    deleted_at TIMESTAMP DEFAULT NULL,
		    FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE
		)
	`
)
