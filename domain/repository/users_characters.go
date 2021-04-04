package repository

import (
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
)

type IUsersCharactersRepository interface {
	Insert(user *domain.User, character *domain.Character) error
	FindByUser(user *domain.User) (*[]*domain.Character, *[]*int, error)
}
