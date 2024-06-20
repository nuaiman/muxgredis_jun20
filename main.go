package main

import (
	"fmt"
	"muxgres/db"
	"muxgres/routes"
	"net/http"
)

func main() {
	db.Init()

	router := http.NewServeMux()

	routes.RegisterRoutes(router)

	server := http.Server{
		Addr:    ":2568",
		Handler: router,
	}
	fmt.Println("Server started at :2568")
	server.ListenAndServe()
}
