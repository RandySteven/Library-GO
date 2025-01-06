package mysql_client

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/enums"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
	"math/rand"
	"time"
)

const (
	defaultPass = `test_1234`
)

func (c *MySQLClient) SeedUserData(ctx context.Context) error {
	time := time.Now()
	users := []*models.User{
		{Name: "Alice", Email: "alice@gmail.com", Password: utils.HashPassword(defaultPass), Address: "John Doe Street", PhoneNumber: "+628123456780", DoB: time},
		{Name: "Brandon", Email: "brandon@gmail.com", Password: utils.HashPassword(defaultPass), Address: "John Doe Street", PhoneNumber: "+628123456781", DoB: time},
		{Name: "Charlie", Email: "charlie@gmail.com", Password: utils.HashPassword(defaultPass), Address: "John Doe Street", PhoneNumber: "+628123456782", DoB: time},
		{Name: "David", Email: "david@gmail.com", Password: utils.HashPassword(defaultPass), Address: "John Doe Street", PhoneNumber: "+628123456783", DoB: time},
		{Name: "Eric", Email: "eric@gmail.com", Password: utils.HashPassword(defaultPass), Address: "John Doe Street", PhoneNumber: "+628123456784", DoB: time},
		{Name: "Fin", Email: "fin@gmail.com", Password: utils.HashPassword(defaultPass), Address: "John Doe Street", PhoneNumber: "+628123456785", DoB: time},
	}
	for _, user := range users {
		_, err := c.db.ExecContext(ctx, queries.InsertUserQuery.ToString(), &user.Name, &user.Address, &user.Email, &user.PhoneNumber, &user.Password, &user.DoB)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *MySQLClient) SeedAuthorData(ctx context.Context) error {
	authors := []*models.Author{
		{Name: "Haruki Murakami", Nationality: "Japanese"},
		{Name: "Chimamanda Ngozi Adichie", Nationality: "Nigerian"},
		{Name: "Gabriel García Márquez", Nationality: "Colombian"},
		{Name: "J.K. Rowling", Nationality: "British"},
		{Name: "Toni Morrison", Nationality: "American"},
	}
	for _, author := range authors {
		_, err := c.db.ExecContext(ctx, queries.InsertAuthorQuery.ToString(), &author.Name, &author.Nationality)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *MySQLClient) SeedBookData(ctx context.Context) error {
	books := []*models.Book{
		{Title: "Book Title A", Description: "Book description A", Image: "", Status: enums.Available},
		{Title: "Book Title B", Description: "Book description B", Image: "", Status: enums.Available},
		{Title: "Book Title C", Description: "Book description C", Image: "", Status: enums.Available},
		{Title: "Book Title D", Description: "Book description D", Image: "", Status: enums.Available},
		{Title: "Book Title E", Description: "Book description E", Image: "", Status: enums.Available},
		{Title: "Book Title F", Description: "Book description F", Image: "", Status: enums.Available},
		{Title: "Book Title G", Description: "Book description G", Image: "", Status: enums.Available},
		{Title: "Book Title H", Description: "Book description H", Image: "", Status: enums.Available},
		{Title: "Book Title I", Description: "Book description I", Image: "", Status: enums.Available},
		{Title: "Book Title J", Description: "Book description J", Image: "", Status: enums.Available},
		{Title: "Book Title K", Description: "Book description K", Image: "", Status: enums.Available},
		{Title: "Book Title L", Description: "Book description L", Image: "", Status: enums.Available},
	}

	for _, book := range books {
		result, err := c.db.ExecContext(ctx, queries.InsertBookQuery.ToString(), &book.Title, &book.Description, &book.Image, &book.Status)
		if err != nil {
			return err
		}
		id, _ := result.LastInsertId()
		for i := 0; i < 2; i++ {
			rand.Seed(time.Now().UnixNano())

			randomNumber := rand.Intn(5) + 1
			_, err = c.db.ExecContext(ctx, queries.InsertAuthorBookQuery.ToString(), randomNumber, id)
			if err != nil {
				return err
			}
		}
		for i := 0; i < 2; i++ {
			rand.Seed(time.Now().UnixNano())

			randomNumber := rand.Intn(53) + 1

			_, err = c.db.ExecContext(ctx, queries.InsertBookGenreQuery.ToString(), id, randomNumber)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func (c *MySQLClient) SeedRoleData(ctx context.Context) error {
	roles := []*models.Role{
		{ID: uint64(enums.Admin), Role: `Admin`},
		{ID: uint64(enums.Member), Role: `Member`},
		{ID: uint64(enums.Premium), Role: `Premium`},
		{ID: uint64(enums.Librarian), Role: `Librarian`},
	}

	query := `INSERT INTO roles (id, role) VALUES (?, ?)`

	for _, role := range roles {
		_, err := c.db.ExecContext(ctx, query, role.ID, role.Role)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *MySQLClient) SeedRoleUserData(ctx context.Context) error {
	roleUsers := []*models.RoleUser{
		{UserID: 1, RoleID: uint64(enums.Admin)},
		{UserID: 2, RoleID: uint64(enums.Librarian)},
		{UserID: 3, RoleID: uint64(enums.Member)},
		{UserID: 4, RoleID: uint64(enums.Member)},
		{UserID: 5, RoleID: uint64(enums.Member)},
		{UserID: 6, RoleID: uint64(enums.Member)},
	}

	query := `INSERT INTO role_users (user_id, role_id) VALUES (?, ?)`

	for _, roleUser := range roleUsers {
		_, err := c.db.ExecContext(ctx, query, roleUser.UserID, roleUser.RoleID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *MySQLClient) SeedGenreData(ctx context.Context) error {
	genres := []*models.Genre{
		{Genre: `Adventure`},
		{Genre: `Mystery`},
		{Genre: `Thriller`},
		{Genre: `Fantasy`},
		{Genre: `Historical Fiction`},
		{Genre: `Science Fiction`},
		{Genre: `Horror`},
		{Genre: `Dystopian`},
		{Genre: `Magical Realism`},
		{Genre: `Literary Fiction`},
		{Genre: `Biography`},
		{Genre: `Memoir`},
		{Genre: `Self-Help`},
		{Genre: `True Crime`},
		{Genre: `Essay`},
		{Genre: `History`},
		{Genre: `Science`},
		{Genre: `Philosophy`},
		{Genre: `Psychology`},
		{Genre: `Travel`},
		{Genre: `Picture Books`},
		{Genre: `Early Readers`},
		{Genre: `Middle Grade`},
		{Genre: `Young Adult`},
		{Genre: `Narrative Poetry`},
		{Genre: `Lyrical Poetry`},
		{Genre: `Haiku`},
		{Genre: `Epic Poetry`},
		{Genre: `Tragedy`},
		{Genre: `Comedy`},
		{Genre: `Historical Drama`},
		{Genre: `Modern Drama`},
		{Genre: `Contemporary Romance`},
		{Genre: `Historical Romance`},
		{Genre: `Paranormal Romance`},
		{Genre: `Romantic Suspense`},
		{Genre: `Dark Fantasy`},
		{Genre: `Urban Fantasy`},
		{Genre: `High Fantasy`},
		{Genre: `Sword and Sorcery`},
		{Genre: `Space Opera`},
		{Genre: `Cyberpunk`},
		{Genre: `Hard Sci-Fi`},
		{Genre: `Historical Romance`},
		{Genre: `Historical Thriller`},
		{Genre: `Alternate History`},
		{Genre: `Theology`},
		{Genre: `Devotionals`},
		{Genre: `Spiritual Growth`},
		{Genre: `Superhero`},
		{Genre: `Manga`},
		{Genre: `Slice of Life`},
		{Genre: `Fantasy Graphic Novels`},
	}
	for _, genre := range genres {
		_, err := c.db.ExecContext(ctx, queries.InsertGenreQuery.ToString(), genre.Genre)
		if err != nil {
			return err
		}
	}
	return nil
}
