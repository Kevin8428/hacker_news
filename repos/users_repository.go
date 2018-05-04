package repos

import "github.com/kevin8428/hackernews/domain"

type UsersRepositoryInterface interface {
	FindUsersByUserID(int) domain.User
}
