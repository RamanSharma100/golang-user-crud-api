package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, status int, error interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func ParseBody(r *http.Request, data interface{}) error {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), data); err != nil {
			return err
		}
	}
	return nil
}
