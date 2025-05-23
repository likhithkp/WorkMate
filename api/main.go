package main

import (
	"net/http"

	"github.com/likhithkp/WorkMate/api/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.UserRouter(mux)
	http.ListenAndServe(":3001", mux)
}
