package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Error(w http.ResponseWriter, statusCode int) error {
	type Error struct {
		Status     string `json:"status"`
		StatusCode int    `json:"status_code"`
	}

	res := struct {
		Error Error `json:"error"`
	}{
		Error: Error{
			Status:     http.StatusText(statusCode),
			StatusCode: statusCode,
		},
	}

	w.WriteHeader(statusCode)
	enc := json.NewEncoder(w)
	err := enc.Encode(&res)
	if err != nil {
		return fmt.Errorf("HTTP Response Error: %w", err)
	}

	return nil
}
