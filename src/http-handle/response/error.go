package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func Error(w http.ResponseWriter, err error) error {
	type Error struct{}

	res := struct {
		Error Error `json:"error"`
	}{
		Error: Error{},
	}

	writeHeader(w)
	w.WriteHeader(http.StatusInternalServerError) // TODO
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		return fmt.Errorf("http-handle/response.Error: %w", err)
	}

	return nil
}
