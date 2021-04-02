package repository

import (
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
)

type ICharacterRepository interface {
	BindByID(id *domain.CharacterID) (*domain.Character, error)
	GetMaxID() (*domain.CharacterID, error)
}
