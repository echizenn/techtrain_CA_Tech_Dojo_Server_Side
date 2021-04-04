package application

import (
	"fmt"
	"strconv"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository"
	"golang.org/x/xerrors"
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
	UserCharacterID string `json:"userCharacterID"`
	CharacterID     string `json:"characterID"`
	Name            string `json:"name"`
}
type UserHoldCharacters []UserHoldCharacter

func (hcas *UsersCharactersApplicationService) Hold(token string) (*UserHoldCharacters, error) {
	targetToken, err := domain.NewUserToken(token)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	user, err := hcas.userRepository.FindByToken(targetToken)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	characters, ids, err := hcas.usersCharactersRepository.FindByUser(user)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	for _, id := range *ids {
		fmt.Printf(strconv.Itoa(*id))
	}

	var userHoldCharacters UserHoldCharacters

	for i := range *characters {
		character := (*characters)[i]
		id := (*ids)[i]

		var userHoldCharacter UserHoldCharacter
		userHoldCharacter.UserCharacterID = strconv.Itoa(*id)
		intCharacterID := character.GetID()
		userHoldCharacter.CharacterID = strconv.Itoa(int(intCharacterID))
		userHoldCharacter.Name = string(character.GetName())
		userHoldCharacters = append(userHoldCharacters, userHoldCharacter)
	}

	return &userHoldCharacters, nil
}
