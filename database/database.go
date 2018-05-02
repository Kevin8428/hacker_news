package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "kdeutscher"
	Password = foo
	DBname   = "news_app"
)

// DatabaseHandler something
type DatabaseHandler struct {
	DB *sql.DB
}

func (h *DatabaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var name string
	// Execute the query.
	row := h.DB.QueryRow("SELECT * FROM users")
	if err := row.Scan(&name); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Write it back to the client.
	fmt.Fprintf(w, "Hi, %s\n", name)
}

// GetUserFirstArticle function
func (h *DatabaseHandler) GetUserFirstArticle(userID int) string {

	rows, err := h.DB.Query("select name from user_articles where user_id = $1 LIMIT 1;", userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var article string
	for rows.Next() {
		if err := rows.Scan(&article); err != nil {
			log.Fatal(err)
		}
	}
	return article
}

// GetUserInfo function
func (h *DatabaseHandler) GetUserLastName(userID int) string {

	rows, err := h.DB.Query("SELECT last_name FROM users WHERE id = $1", userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var lastName string
	for rows.Next() {
		if err := rows.Scan(&lastName); err != nil {
			log.Fatal(err)
		}
	}
	return lastName
}

// AddArticle does stuff
func (h *DatabaseHandler) AddArticle(name string, author string, website string, id int) {

	_, err := h.DB.Query("INSERT INTO user_articles (name, author, website, user_id) VALUES ($1, $2, $3, $4)", name, author, website, id)
	if err != nil {
		fmt.Println("insert error: ", err)
		panic(err)
	}
}
