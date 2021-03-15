package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/application"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/service"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/infrastructure"
)

type gachaDrawJson struct {
	Times int `json:"times"`
}

func (api *GameAPI) GachaDraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("POSTだけです。"))
		return
	}

	header := r.Header
	token := header["X-Token"][0] // なんで大文字になる？、0って明示して大丈夫？

	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, body)

	var gdj gachaDrawJson
	json.Unmarshal(buf.Bytes(), &gdj)

	times := gdj.Times

	ur := infrastructure.NewUserRepository(api.db)
	cr := infrastructure.NewCharacterRepository(api.db)
	ucr := infrastructure.NewUsersCharactersRepository(cr, api.db)

	gs := service.NewGachaService(cr)

	gas := application.NewGachaApplicationService(ur, ucr, gs)

	var results []application.GachaDrawResult

	for i := 0; i < times; i++ {
		gachaDrawResult, err := gas.Draw(token)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, *gachaDrawResult)
	}

	stringResults, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("results", string(stringResults))
	w.WriteHeader(http.StatusOK)
}