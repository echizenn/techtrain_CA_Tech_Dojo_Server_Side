package application

import (
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository_interface"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/service"
)

type UserApplicationService struct {
	userRepository   repository_interface.IUserRepository
	userIdService    service.UserIdService
	userTokenService service.UserTokenService
}

func NewUserApplicationService(userRepository repository_interface.IUserRepository,
	userIdService service.UserIdService,
	userTokenService service.UserTokenService) UserApplicationService {
	return UserApplicationService{userRepository, userIdService, userTokenService}
}

func (uas UserApplicationService) Register(name string) (*string, error) {
	userId := uas.userIdService.Create()

	userName, err := domain.NewName(name)
	if err != nil {
		return nil, err
	}

	userToken := uas.userTokenService.Create()

	user := domain.NewUser(*userId, *userName, *userToken)
	err = uas.userRepository.Insert(&user)
	if err != nil {
		return nil, err
	}
	stringUserToken := string(user.GetToken())
	return &stringUserToken, nil
}

func (uas UserApplicationService) GetName(token string) (*string, error) {
	targetToken, err := domain.NewToken(token)
	if err != nil {
		return nil, err
	}

	user, err := uas.userRepository.FindByToken(targetToken)
	if err != nil {
		return nil, err
	}

	name := user.GetName()
	stringName := string(name)
	return &stringName, nil
}

func (uas UserApplicationService) Update(name string, token string) error {
	userToken, err := domain.NewToken(token)
	if err != nil {
		return err
	}
	user, err := uas.userRepository.FindByToken(userToken)
	if err != nil {
		return err
	}

	newName, err := domain.NewName(name)
	if err != nil {
		return err
	}

	newUser := user.SetName(*newName)

	err = uas.userRepository.Update(&newUser)
	if err != nil {
		return err
	}
	return nil
}
