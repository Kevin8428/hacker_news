package articles

import (
	"github.com/kevin8428/hackernews/repos"
)

type Service interface {
	GetArticleInfo(string) string
}

type service struct {
	Articles repos.ArticlesRepositoryInterface
}

// NewService comment
func NewService(articleRepo repos.ArticlesRepositoryInterface) Service {
	return &service{
		Articles: articleRepo,
	}
}

func (s *service) GetArticleInfo(author string) string {
	s.Articles.FindArticlesByUserID(1)
	return "author name"
}
