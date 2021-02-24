package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/application"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/repository"
)

func UserHoldCharacterList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		// ここの処理怪しさしかない
		w.Write([]byte("GETだけです。"))
		return
	}

	header := r.Header
	token := header["X-Token"][0] // なんで大文字になる？、0って明示して大丈夫？

	ur := repository.NewUserRepository()
	cr := repository.NewCharacterRepository()
	ucr := repository.NewUsersCharactersRepository(cr)

	ucas := application.NewUsersCharactersApplicationService(ur, ucr)

	userHoldCharacters, err := ucas.Hold(token)
	if err != nil {
		log.Fatal(err)
	}

	stringCharacters, err := json.Marshal(userHoldCharacters)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("characters", string(stringCharacters))
	w.WriteHeader(http.StatusOK)
}
