package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

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

	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, body)

	var gdj gachaDrawJson
	json.Unmarshal(buf.Bytes(), &gdj)

	times := gdj.Times

	var results []application.GachaDrawResult

	for i := 0; i < times; i++ {
		gachaDrawResult, err := api.gas.Draw(token)
		if err != nil {
			return err
		}

		results = append(results, *gachaDrawResult)
	}

	stringResults, err := json.Marshal(results)
	if err != nil {
		return err
	}

	w.Header().Set("results", string(stringResults))
	w.WriteHeader(http.StatusOK)

	return nil
}

func (api *GameAPI) GachaDrawHandler(w http.ResponseWriter, r *http.Request) {
	err := api.GachaDraw(w, r)
	if err != nil {
		log.Fatal(err)
	}
}
