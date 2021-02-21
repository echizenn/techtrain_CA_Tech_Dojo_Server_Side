package repository_interface

import (
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
)

type IUserRepository interface {
	Save(user *domain.User)
	FindByToken(token *domain.UserToken)
}
