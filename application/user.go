package application

import (
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/service"
	"golang.org/x/xerrors"
)

type UserApplicationService struct {
	userRepository   repository.IUserRepository
	userIDService    service.UserIDService
	userTokenService service.UserTokenService
}

func NewUserApplicationService(userRepository repository.IUserRepository,
	userIDService service.UserIDService,
	userTokenService service.UserTokenService) UserApplicationService {
	return UserApplicationService{userRepository, userIDService, userTokenService}
}

func (userApplicationService *UserApplicationService) Register(name string) (*string, error) {
	userID, err := userApplicationService.userIDService.Create()
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	userName, err := domain.NewUserName(name)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	userToken, err := userApplicationService.userTokenService.Create()
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	user := domain.NewUser(*userID, *userName, *userToken)
	err = userApplicationService.userRepository.Insert(&user)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}
	stringUserToken := string(user.GetToken())
	return &stringUserToken, nil
}

func (userApplicationService *UserApplicationService) GetName(token string) (*string, error) {
	targetToken, err := domain.NewUserToken(token)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	user, err := userApplicationService.userRepository.FindByToken(targetToken)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	name := user.GetName()
	stringName := string(name)
	return &stringName, nil
}

func (userApplicationService *UserApplicationService) Update(name string, token string) error {
	userToken, err := domain.NewUserToken(token)
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}
	user, err := userApplicationService.userRepository.FindByToken(userToken)
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	newName, err := domain.NewUserName(name)
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	newUser := user.SetName(*newName)

	err = userApplicationService.userRepository.Update(newUser)
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}
	return nil
}
