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
	query := `SELECT last_name FROM users WHERE id = $1`
	rows, error := u.db.Query(query, userID)
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
