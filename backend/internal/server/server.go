package server

import (
	"net/http"
)

func Server(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/build/index.html")

}