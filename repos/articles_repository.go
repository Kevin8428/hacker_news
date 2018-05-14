package repos

import "github.com/kevin8428/hackernews/domain"

type ArticlesRepositoryInterface interface {
	FindArticlesByUserID(int) []domain.Article
	FindSportsArticles() []domain.Article
}
