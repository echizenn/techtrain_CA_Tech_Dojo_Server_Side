package application

import (
	"strconv"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/service"
	"golang.org/x/xerrors"
)

type GachaApplicationService struct {
	userRepository            repository.IUserRepository
	usersCharactersRepository repository.IUsersCharactersRepository
	gachaService              service.GachaService
}

func NewGachaApplicationService(
	userRepository repository.IUserRepository,
	usersCharactersRepository repository.IUsersCharactersRepository,
	gachaService service.GachaService) GachaApplicationService {
	return GachaApplicationService{userRepository, usersCharactersRepository, gachaService}
}

type GachaDrawResult struct {
	CharacterID string `json:"characterID"`
	Name        string `json:"name"`
}

func (gachaApplicationService *GachaApplicationService) Draw(token string) (*GachaDrawResult, error) {
	targetToken, err := domain.NewUserToken(token)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	user, err := gachaApplicationService.userRepository.FindByToken(targetToken)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	character, err := gachaApplicationService.gachaService.Draw()
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	err = gachaApplicationService.usersCharactersRepository.Insert(user, character)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	characterID := character.GetID()
	intCharacterID := int(characterID)
	stringCharacterID := strconv.Itoa(intCharacterID)

	characterName := character.GetName()
	stringCharacterName := string(characterName)

	return &GachaDrawResult{stringCharacterID, stringCharacterName}, nil
}
