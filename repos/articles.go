package repos

import (
	"database/sql"

	"github.com/kevin8428/hackernews/api"
)

type ArticlesRepository struct {
	db *sql.DB
}

func (a *ArticlesRepository) FindArticlesByUserID(int) []api.Article {
	return []api.Article{}
}
