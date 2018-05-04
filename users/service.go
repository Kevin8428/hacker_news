package users

import (
	"github.com/kevin8428/hackernews/domain"
	"github.com/kevin8428/hackernews/repos"
)

type Service interface {
	FindUsersByUserID(int) domain.User
}

type service struct {
	Users repos.UsersRepositoryInterface
}

// NewService comment
func NewService(userRepo repos.UsersRepositoryInterface) Service {
	return &service{
		Users: userRepo,
	}
}

func (s *service) FindUsersByUserID(id int) domain.User {
	s.Users.FindUsersByUserID(id)
	return domain.User{}
}
