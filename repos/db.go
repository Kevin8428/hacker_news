package repos

import (
	"database/sql"
	"fmt"

	"github.com/kevin8428/hackernews/database"
)

type Repositories struct {
	Articles *ArticlesRepository
	Users    *UsersRepository
}

func Initialize() *Repositories {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", database.Host, database.Port, database.User, database.Password, database.DBname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error connecting to DB: ", err)
	}
	return &Repositories{
		Articles: &ArticlesRepository{db},
		Users:    &UsersRepository{db},
	}
}
