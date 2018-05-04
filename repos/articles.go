package repos

import (
	"database/sql"

	"github.com/kevin8428/hackernews/domain"
)

type ArticlesRepository struct {
	DB *sql.DB
}

func (a *ArticlesRepository) FindArticlesByUserID(int) []domain.Article {
	return []domain.Article{}
}
