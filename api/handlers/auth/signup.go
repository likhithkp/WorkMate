package handlers

import (
	"bytes"
	"io"
	"net/http"

	"github.com/likhithkp/WorkMate/api/helpers"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.WriteJson(w, "", 405, "Method not allowed")
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.WriteJson(w, "", 500, "Failed to read the body stream")
		return
	}
	defer r.Body.Close()

	endpointURL := "http://localhost:3000/signup"
	req, err := http.NewRequest("POST", endpointURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to forward request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
