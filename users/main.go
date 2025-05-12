package main

import (
	"net/http"

	"github.com/likhithkp/social-media-backend/users/routes"
)

func main() {
	mux := http.NewServeMux()

	routes.UserRouter(mux)
	http.ListenAndServe(":3000", mux)
}
