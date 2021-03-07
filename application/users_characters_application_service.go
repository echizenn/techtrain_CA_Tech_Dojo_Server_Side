package application

import (
	"strconv"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository"
)

type UsersCharactersApplicationService struct {
	userRepository            repository.IUserRepository
	usersCharactersRepository repository.IUsersCharactersRepository
}

func NewUsersCharactersApplicationService(
	userRepository repository.IUserRepository,
	usersCharactersRepository repository.IUsersCharactersRepository) UsersCharactersApplicationService {
	return UsersCharactersApplicationService{userRepository, usersCharactersRepository}
}

type UserHoldCharacter struct {
	// userCharacterId string `json:"userCharacterID"`
	CharacterId string `json:"characterID"`
	Name        string `json:"name"`
}
type UserHoldCharacters []UserHoldCharacter

func (hcas *UsersCharactersApplicationService) Hold(token string) (*UserHoldCharacters, error) {
	targetToken, err := domain.NewUserToken(token)
	if err != nil {
		return nil, err
	}

	user, err := hcas.userRepository.FindByToken(targetToken)
	if err != nil {
		return nil, err
	}

	characters, err := hcas.usersCharactersRepository.FindByUser(user)
	if err != nil {
		return nil, err
	}

	var userHoldCharacters UserHoldCharacters

	for _, character := range *characters {
		var userHoldCharacter UserHoldCharacter
		intCharacterId := character.GetId()
		userHoldCharacter.CharacterId = strconv.Itoa(int(intCharacterId))
		userHoldCharacter.Name = string(character.GetName())
		userHoldCharacters = append(userHoldCharacters, userHoldCharacter)
	}

	return &userHoldCharacters, nil
}
