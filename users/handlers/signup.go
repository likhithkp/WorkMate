package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/likhithkp/social-media-backend/users/shared"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		res := shared.Response{
			Message:    "Not a valid method",
			StatusCode: 405,
			Data:       nil,
		}

		if err := json.NewEncoder(w).Encode(&res); err != nil {
			log.Println("Failed to encode the response for `Not a valid method`")
			return
		}
		return
	}

	newUser := new(shared.SignUp)

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		res := shared.Response{
			Message:    "Error while decoding the signup body",
			StatusCode: 500,
			Data:       nil,
		}

		if err := json.NewEncoder(w).Encode(&res); err != nil {
			log.Println("Failed to encode the error message `Error while decoding the signup body`")
			return
		}
		return
	}

	if newUser.Age < 18 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		res := shared.Response{
			Message:    "Sorry, you must be atleast 18 year old to have an account!",
			StatusCode: 400,
			Data:       nil,
		}

		if err := json.NewEncoder(w).Encode(&res); err != nil {
			log.Println("Failed to encode the error message `Sorry, you must be atleast 18 year old to have an account!`")
			return
		}
		return
	}

	if newUser.Firstname == "" || newUser.Surname == "" || newUser.Age == 0 || newUser.Gender == "" || newUser.MobileNumber == 0 || newUser.NewPassword == "" || newUser.Email == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		res := shared.Response{
			Message:    "Invalid value / Missing fields",
			StatusCode: 400,
			Data:       nil,
		}

		if err := json.NewEncoder(w).Encode(&res); err != nil {
			log.Println("Failed to encode the error message `Invalid value / Missing fields`")
			return
		}
		return
	}
}
