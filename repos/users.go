package repos

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kevin8428/hackernews/domain"
)

// UsersRepository struct
type UsersRepository struct {
	DB *sql.DB
}

// FindUsersByUserID method
func (u *UsersRepository) FindUsersByUserID(userID int) domain.User {
	query := `SELECT last_name FROM users WHERE id = $1`
	rows, error := u.DB.Query(query, userID)
	if error != nil {
		fmt.Println("error1: ", error)
		panic(error)
		return domain.User{}
	}
	defer rows.Close()
	var lastName string
	for rows.Next() {
		if err := rows.Scan(&lastName); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("lastName: ", lastName)
	return domain.User{
		LastName: lastName,
	}
}

// SaveArticle comment
func (u *UsersRepository) SaveArticle(name string, author string, website string, id int) error {
	_, err := u.DB.Query("INSERT INTO user_articles (name, author, website, user_id) VALUES ($1, $2, $3, $4)", name, author, website, id)
	if err != nil {
		fmt.Println("insert error: ", err)
		panic(err)
	}
	return err
}
