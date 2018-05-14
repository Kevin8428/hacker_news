package repos

import (
	"database/sql"
	"fmt"

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
	stmt, err := u.DB.Prepare("SELECT id, first_name, last_name, email, password, auth_token FROM users WHERE id = $1")
	var (
		ID        int
		FirstName string
		LastName  string
		Email     string
		Password  string
		AuthToken string
	)
	err = stmt.QueryRow(userID).Scan(&ID, &FirstName, &LastName, &Email, &Password, &AuthToken)
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
		AuthToken:  AuthToken,
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
	var password sql.NullString
	var token sql.NullString
	sqlStatement := "SELECT password, auth_token FROM users WHERE email = $1"
	row := u.DB.QueryRow(sqlStatement, email)
	switch err := row.Scan(&password, &token); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return "", "", sql.ErrNoRows
	case nil:
		fmt.Println("query worked: ", password, token)
		return password.String, token.String, nil
	default:
		panic(err)
	}
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
