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
	CharacterId string `json:"characterID"`
	Name        string `json:"name"`
}

func (gas *GachaApplicationService) Draw(token string) (*GachaDrawResult, error) {
	targetToken, err := domain.NewUserToken(token)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	user, err := gas.userRepository.FindByToken(targetToken)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	character, err := gas.gachaService.Draw()
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	err = gas.usersCharactersRepository.Insert(user, character)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	characterId := character.GetId()
	intCharacterId := int(characterId)
	stringCharacterId := strconv.Itoa(intCharacterId)

	characterName := character.GetName()
	stringCharacterName := string(characterName)

	return &GachaDrawResult{stringCharacterId, stringCharacterName}, nil
}
