package repository

import (
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
)

type IUserRepository interface {
	Insert(user *domain.User) error
	Update(user *domain.User) error
	FindByToken(token *domain.UserToken) (*domain.User, error)
	GetMaxID() (*domain.UserID, error)
}
