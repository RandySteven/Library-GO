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
)
