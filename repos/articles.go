package repos

import (
	"database/sql"
	"fmt"

	"github.com/kevin8428/hackernews/domain"
)

type ArticlesRepository struct {
	DB *sql.DB
}

func (a *ArticlesRepository) FindArticlesByUserID(int) []domain.Article {
	return []domain.Article{}
}

func (a *ArticlesRepository) FindSportsArticles() []domain.Article {
	articles := []domain.Article{}
	rows, err := a.DB.Query("SELECT title, website, link FROM sports_articles")
	if err != nil {
		fmt.Println("cant find record: ", err)
		return []domain.Article{}
	}
	defer rows.Close()
	for rows.Next() {
		var (
			title   sql.NullString
			website sql.NullString
			url     sql.NullString
		)
		if err := rows.Scan(&title, &website, &url); err != nil {
			fmt.Printf("error scanning articles: %v", err)
			continue
		}
		article := domain.Article{
			Title:   title.String,
			Website: website.String,
			URL:     url.String,
		}
		articles = append(articles, article)
	}
	return articles
}
