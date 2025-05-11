package routes

import (
	"net/http"

	"github.com/likhithkp/social-media-backend/users/handlers"
)

func UserRouter(mux *http.ServeMux) {
	mux.HandleFunc("/signup", handlers.SignUp)
}
