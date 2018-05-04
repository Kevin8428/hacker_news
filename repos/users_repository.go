package repos

import "github.com/kevin8428/hackernews/domain"

type UsersRepositoryInterface interface {
	FindUsersByUserID(int) domain.User
	SaveArticle(string, string, string, int, string, string) error
	FindUserArticlesByUserID(int) []domain.Article
	FindUserByAuthToken(string) domain.User
}
