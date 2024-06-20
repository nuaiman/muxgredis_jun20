package routes

import (
	"net/http"
)

func RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/", handleRoot)
}
