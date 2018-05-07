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

func (u *UsersRepository) FindUserByAuthToken(token string) domain.User {
	stmt, err := u.DB.Prepare("SELECT id, first_name, last_name, email, password, auth_token FROM users WHERE auth_token = $1")
	var (
		ID        int
		FirstName string
		LastName  string
		Email     string
		Password  string
		AuthToken string
	)
	err = stmt.QueryRow(token).Scan(&ID, &FirstName, &LastName, &Email, &Password, &AuthToken)
	if err != nil {
		fmt.Println("cant find record: ", err)
		return domain.User{}
	}
	isLoggedIn := ID > 0
	return domain.User{
		ID:         ID,
		FirstName:  FirstName,
		LastName:   LastName,
		Password:   Password,
		IsLoggedIn: isLoggedIn,
	}
}

// FindUserArticlesByUserID is a method
func (u *UsersRepository) FindUserArticlesByUserID(id int) []domain.Article {
	articles := []domain.Article{}
	rows, err := u.DB.Query("SELECT user_id, title, website, category, author, link FROM user_articles WHERE user_id = $1", id)
	if err != nil {
		fmt.Println("cant find record: ", err)
		return []domain.Article{}
	}
	defer rows.Close()
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

func (u *UsersRepository) GetPasswordUsingEmail(email string) (string, string, error) {
	rows, err := u.DB.Query("SELECT password, auth_token FROM users WHERE email = $1", email)
	if err != nil {
		fmt.Println("cant find record: ", err)
		return "", "", err
	}
	defer rows.Close()
	var password sql.NullString
	var token sql.NullString
	for rows.Next() {
		if err := rows.Scan(&password, &token); err != nil {
			fmt.Printf("error scanning: %v", err)
			return "", "", err
		}
	}
	return password.String, token.String, nil
}

// CreateUser(email, password, fn, ln)
func (u *UsersRepository) CreateUser(token string, email string, password string, fn string, ln string) (string, error) {
	_, err := u.DB.Query("INSERT INTO users (first_name, last_name, email, password, auth_token) VALUES ($1, $2, $3, $4, $5)", fn, ln, email, password, token)
	if err != nil {
		fmt.Println("insert error: ", err)
		return "", err
	}
	return token, err
}
