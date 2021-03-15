package api

import (
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/application"
	_ "github.com/go-sql-driver/mysql"
)

type GameAPI struct {
	uas  application.UserApplicationService
	gas  application.GachaApplicationService
	ucas application.UsersCharactersApplicationService
}

func NewGameAPI(uas application.UserApplicationService,
	gas application.GachaApplicationService,
	ucas application.UsersCharactersApplicationService,
) GameAPI {
	return GameAPI{uas, gas, ucas}
}
