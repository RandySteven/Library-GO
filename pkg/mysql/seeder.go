package mysql

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
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
