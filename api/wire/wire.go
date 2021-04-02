//+build wireinject

package wire

import (
	"database/sql"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/api"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/application"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/service"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/infrastructure"
	"github.com/google/wire"
)

func InitGameAPI(db *sql.DB) api.GameAPI {
	wire.Build(
		infrastructure.NewUserRepository,
		infrastructure.NewCharacterRepository,
		infrastructure.NewUsersCharactersRepository,
		service.NewGachaService,
		service.NewUserIDService,
		service.NewUserTokenService,
		application.NewUserApplicationService,
		application.NewGachaApplicationService,
		application.NewUsersCharactersApplicationService,
		api.NewGameAPI)

	return api.GameAPI{}
}
