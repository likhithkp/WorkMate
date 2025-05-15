package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/likhithkp/WorkMate/api/shared"
)

func WriteJson(w http.ResponseWriter, data any, status uint16, message string) {
	w.Header().Set("Content-Type", "application/json")

	response := shared.Response{
		Message:    message,
		StatusCode: status,
		Data:       data,
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(&response); err != nil {
		http.Error(w, "Internal Server Error: Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(int(status))
	_, _ = w.Write(buf.Bytes())
}
