package repository

import (
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
)

type ICharacterRepository interface {
	FindById(id *domain.CharacterId) (*domain.Character, error)
	GetMaxId() (*domain.CharacterId, error)
}
