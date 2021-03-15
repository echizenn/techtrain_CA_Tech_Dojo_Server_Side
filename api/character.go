package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (api *GameAPI) UserHoldCharacterList(w http.ResponseWriter, r *http.Request) {
	// 確認が重複になるのでいらない気もする
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("GETだけです。"))
		return
	}

	header := r.Header
	token := header["X-Token"][0] // なんで大文字になる？、0って明示して大丈夫？

	userHoldCharacters, err := api.ucas.Hold(token)
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
