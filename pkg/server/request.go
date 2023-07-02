package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func ParseParams(r *http.Request) map[string]string {
	params := mux.Vars(r)
	return params
}

func ParseBody(r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return err
	}
	return nil
}
