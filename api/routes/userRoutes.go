package routes

import (
	"net/http"

	"github.com/likhithkp/WorkMate/api/handlers"
)

func UserRouter(mux *http.ServeMux) {
	mux.HandleFunc("/signup", handlers.SignUp)
}
