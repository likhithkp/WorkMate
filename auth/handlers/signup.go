package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/likhithkp/WorkMate/auth/clients/kafka"
	"github.com/likhithkp/WorkMate/auth/helpers"
	"github.com/likhithkp/WorkMate/auth/shared"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	newUser := new(shared.SignUp)

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		helpers.WriteJson(w, "", 500, "Error while decoding the signup body")
		return
	}

	if newUser.Age < 18 {
		helpers.WriteJson(w, "", 400, "Sorry, you must be atleast 18 year old to have an account!")
		return
	}

	if newUser.Firstname == "" || newUser.Age == 0 || newUser.Gender == "" || newUser.MobileNumber == 0 || newUser.NewPassword == "" || newUser.Email == "" {
		helpers.WriteJson(w, "", 400, "Invalid value / Missing fields")
		return
	}

	kafka.CreateProducer()

	fmt.Println(newUser)

}
