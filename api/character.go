package api

import (
	"encoding/json"
	"net/http"
)

func (api *GameAPI) UserHoldCharacterList(w http.ResponseWriter, r *http.Request) error {
	// 確認が重複になるのでいらない気もする
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("GETだけです。"))
		return nil
	}

	header := r.Header
	token := header["X-Token"][0] // なんで大文字になる？、0って明示して大丈夫？

	userHoldCharacters, err := api.usersCharactersApplicationService.Hold(token)
	if err != nil {
		return err
	}

	stringCharacters, err := json.Marshal(userHoldCharacters)
	if err != nil {
		return err
	}

	w.Header().Set("characters", string(stringCharacters))

	return nil
}

func (api *GameAPI) UserHoldCharacterListHandler(w http.ResponseWriter, r *http.Request) {
	err := api.UserHoldCharacterList(w, r)
	if err != nil {
		// statusコードを設定
		// ログをはく
		return
	}
	w.WriteHeader(http.StatusOK)
}
