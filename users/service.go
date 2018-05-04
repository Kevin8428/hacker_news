package users

import (
	"github.com/kevin8428/hackernews/domain"
	"github.com/kevin8428/hackernews/repos"
)

// methods of of service struct
type Service interface {
	SaveArticleToUser(string, string, string, int, string, string) error
	SaveNewUser() ServiceResponse
	FindArticles(id int) []domain.Article
	FindUser(id int) domain.User
	FindUserByAuth(token string) (domain.User, error)
}

type service struct {
	Users repos.UsersRepositoryInterface
}

type ServiceResponse struct {
	valid bool
}

// NewService comment
func NewService(userRepo repos.UsersRepositoryInterface) Service {
	return &service{
		Users: userRepo,
	}
}

func (s *service) SaveNewUser() ServiceResponse {
	// create user in DB
	// return list of possible articles
	return ServiceResponse{
		valid: true,
	}
}
func (s *service) FindUserByAuth(token string) (domain.User, error) {
	user := s.Users.FindUserByAuthToken(token)
	return user, nil
}

func (s *service) FindArticles(id int) []domain.Article {
	return s.Users.FindUserArticlesByUserID(id)
}

func (s *service) FindUser(id int) domain.User {
	return s.Users.FindUsersByUserID(id)
}

func (s *service) SaveArticleToUser(name string, author string, website string, id int, category string, url string) error {
	err := s.Users.SaveArticle(name, author, website, id, category, url)
	return err
}
