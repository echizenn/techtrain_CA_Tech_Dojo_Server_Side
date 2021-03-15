package application

import (
	"fmt"
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
	UserCharacterId string `json:"userCharacterID"`
	CharacterId     string `json:"characterID"`
	Name            string `json:"name"`
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

	characters, ids, err := hcas.usersCharactersRepository.FindByUser(user)
	if err != nil {
		return nil, err
	}

	for _, id := range *ids {
		fmt.Printf(strconv.Itoa(*id))
	}

	var userHoldCharacters UserHoldCharacters

	for i := range *characters {
		character := (*characters)[i]
		id := (*ids)[i]

		var userHoldCharacter UserHoldCharacter
		userHoldCharacter.UserCharacterId = strconv.Itoa(*id)
		intCharacterId := character.GetId()
		userHoldCharacter.CharacterId = strconv.Itoa(int(intCharacterId))
		userHoldCharacter.Name = string(character.GetName())
		userHoldCharacters = append(userHoldCharacters, userHoldCharacter)
	}

	return &userHoldCharacters, nil
}
