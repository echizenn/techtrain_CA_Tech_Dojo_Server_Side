package application

import (
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/service"
	"golang.org/x/xerrors"
)

type UserApplicationService struct {
	userRepository   repository.IUserRepository
	userIdService    service.UserIdService
	userTokenService service.UserTokenService
}

func NewUserApplicationService(userRepository repository.IUserRepository,
	userIdService service.UserIdService,
	userTokenService service.UserTokenService) UserApplicationService {
	return UserApplicationService{userRepository, userIdService, userTokenService}
}

func (uas *UserApplicationService) Register(name string) (*string, error) {
	userId, err := uas.userIdService.Create()
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	userName, err := domain.NewUserName(name)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	userToken, err := uas.userTokenService.Create()
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	user := domain.NewUser(*userId, *userName, *userToken)
	err = uas.userRepository.Insert(&user)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}
	stringUserToken := string(user.GetToken())
	return &stringUserToken, nil
}

func (uas *UserApplicationService) GetName(token string) (*string, error) {
	targetToken, err := domain.NewUserToken(token)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	user, err := uas.userRepository.FindByToken(targetToken)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	name := user.GetName()
	stringName := string(name)
	return &stringName, nil
}

func (uas *UserApplicationService) Update(name string, token string) error {
	userToken, err := domain.NewUserToken(token)
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}
	user, err := uas.userRepository.FindByToken(userToken)
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	newName, err := domain.NewUserName(name)
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	newUser := user.SetName(*newName)

	err = uas.userRepository.Update(newUser)
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}
	return nil
}
