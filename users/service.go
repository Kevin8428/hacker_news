package users

import (
	"github.com/kevin8428/hackernews/domain"
	"github.com/kevin8428/hackernews/repos"
)

// methods of of service struct
type Service interface {
	SaveArticle(string, string, string, int) error
	FindUser(int) domain.User
	SaveNewUser() ServiceResponse
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

func (s *service) FindUser(id int) domain.User {
	return s.Users.FindUsersByUserID(id)
}

func (s *service) SaveArticle(name string, author string, website string, id int) error {
	err := s.Users.SaveArticle(name, author, website, id)
	return err
}
