package repos

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kevin8428/hackernews/domain"
)

// UsersRepository struct
type UsersRepository struct {
	db *sql.DB
}

// FindUsersByUserID method
func (u *UsersRepository) FindUsersByUserID(userID int) domain.User {
	fmt.Println("----id: ", userID)
	query := `SELECT last_name FROM users WHERE id = $1`
	// rows, err := u.db.Query("SELECT last_name FROM users WHERE id = $1", userID)
	fmt.Println("-----0----")
	rows, queryError := u.db.Query(query, userID)
	fmt.Println("-----0.1----")
	if queryError != nil {
		fmt.Println("error1: ", queryError)
		return domain.User{}
	}
	fmt.Println("-----1----")
	defer rows.Close()
	var lastName string
	for rows.Next() {
		if err := rows.Scan(&lastName); err != nil {
			log.Fatal(err)
		}
	}
	return domain.User{
		LastName: lastName,
	}
}
