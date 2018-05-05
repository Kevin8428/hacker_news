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
// 	id SERIAL PRIMARY KEY,
// 	first_name TEXT,
// 	last_name TEXT,
// 	email TEXT,
// 	password TEXT,
// 	auth_token TEXT
// );

// CREATE TABLE user_articles (
// 	user_id INT,
// 	title TEXT,
// 	website TEXT,
// 	category TEXT,
// 	author TEXT,
// 	link TEXT
// );

// INSERT INTO users (first_name, last_name, email, password, auth_token) VALUES ('kevin', 'deutscher', 'kevin@mail.com', 'password', '1234');
// INSERT INTO users (first_name, last_name, email, password, auth_token) VALUES ('john', 'doe', 'john@mail.com', 'password', '3242');
// INSERT INTO users (first_name, last_name, email, password, auth_token) VALUES ('joe', 'smith', 'joe@mail.com', 'password', '1212');
