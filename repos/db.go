package repos

import (
	"database/sql"
	"fmt"

	"github.com/kevin8428/hackernews/config"
)

type Repositories struct {
	Articles *ArticlesRepository
	Users    *UsersRepository
}

func Initialize() *Repositories {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error connecting to DB: ", err)
	}
	return &Repositories{
		Articles: &ArticlesRepository{db},
		Users:    &UsersRepository{db},
	}
}

// CREATE TABLE users (
//   id SERIAL PRIMARY KEY,
//   first_name TEXT,
//   last_name TEXT,
//   password TEXT NOT NULL
// );
// CREATE TABLE articles (
//   id SERIAL PRIMARY KEY,
//   title TEXT,
//   site TEXT,
//   category TEXT,
//   author TEXT,
//   link TEXT
// );
// CREATE TABLE users_articles (
//   user_id INT REFERENCES users (id),
//   article_id INT REFERENCES articles (id)
// );
