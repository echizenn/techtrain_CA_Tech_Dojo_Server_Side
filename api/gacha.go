package api

import (
	"encoding/json"
	"net/http"

	"golang.org/x/xerrors"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/application"
)

type gachaDrawJson struct {
	Times int `json:"times"`
}

func (api *GameAPI) GachaDraw(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("POSTだけです。"))
		return nil
	}

	header := r.Header
	token := header["X-Token"][0] // なんで大文字になる？、0って明示して大丈夫？

	var gdj gachaDrawJson
	body := r.Body
	dec := json.NewDecoder(body)
	dec.Decode(&gdj)

	times := gdj.Times

	var results []application.GachaDrawResult

	for i := 0; i < times; i++ {
		gachaDrawResult, err := api.gachaApplicationService.Draw(token)
		if err != nil {
			return xerrors.Errorf("gachaApplicatinService.Draw func error: %w", err)
		}

		results = append(results, *gachaDrawResult)
	}

	stringResults, err := json.Marshal(results)
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	w.Header().Set("results", string(stringResults))

	return nil
}

func (api *GameAPI) GachaDrawHandler(w http.ResponseWriter, r *http.Request) {
	err := api.GachaDraw(w, r)
	if err != nil {
		// statusコードを設定
		// ログをはく
		return
	}
	w.WriteHeader(http.StatusOK)
}
