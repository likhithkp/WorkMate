package main

import (
	"net/http"

	"github.com/likhithkp/WorkMate/auth/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.UserRouter(mux)
	http.ListenAndServe(":3000", mux)
}
