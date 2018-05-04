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

// FindUserArticlesByUserID is a method
func (u *UsersRepository) FindUserArticlesByUserID(id int) []domain.Article {
	articles := []domain.Article{}
	rows, _ := u.DB.Query("SELECT user_id, title, website, category, author, link FROM user_articles WHERE user_id = $1", id)
	for rows.Next() {
		var (
			userID   sql.NullInt64
			title    sql.NullString
			website  sql.NullString
			category sql.NullString
			author   sql.NullString
			url      sql.NullString
		)
		if err := rows.Scan(&userID, &title, &website, &category, &author, &url); err != nil {
			fmt.Printf("error scanning articles: %v", err)
			continue
		}
		article := domain.Article{
			UserID:   int(userID.Int64),
			Title:    title.String,
			Website:  website.String,
			Category: category.String,
			Author:   author.String,
			URL:      url.String,
		}
		articles = append(articles, article)
	}
	return articles
}

// FindUsersByUserID method
func (u *UsersRepository) FindUsersByUserID(userID int) domain.User {
	rows, err := u.DB.Query("SELECT last_name FROM users WHERE id = $1", userID)
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
	return domain.User{
		LastName: lastName,
	}
}

// SaveArticle comment
func (u *UsersRepository) SaveArticle(name string, author string, website string, id int, category string, url string) error {
	_, err := u.DB.Query("INSERT INTO user_articles (title, author, website, user_id, category, link) VALUES ($1, $2, $3, $4, $5, $6)", name, author, website, id, category, url)
	if err != nil {
		fmt.Println("insert error: ", err)
		panic(err)
	}
	return err
}
