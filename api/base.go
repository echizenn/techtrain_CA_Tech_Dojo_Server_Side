package api

import (
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/application"
	_ "github.com/go-sql-driver/mysql"
)

type GameAPI struct {
	userApplicationService            application.UserApplicationService
	gachaApplicationService           application.GachaApplicationService
	usersCharactersApplicationService application.UsersCharactersApplicationService
}

func NewGameAPI(
	userApplicationService application.UserApplicationService,
	gachaApplicationService application.GachaApplicationService,
	usersCharactersApplicationService application.UsersCharactersApplicationService,
) GameAPI {
	return GameAPI{userApplicationService, gachaApplicationService, usersCharactersApplicationService}
}
