package main

import (
	"citybike/pkg/api"
	"citybike/pkg/csv"
	"citybike/pkg/db"
	"fmt"
	"net/http"
	"os"

	_ "citybike/docs"
)

// @title Golang example API
// @version 0.1
// @description This is a simple boilerplate for Go API.

// @contact.name Kimi Porthan

// @host localhost:8080

func main() {
	port := ":" + os.Getenv("APP_PORT")
	// If no PORT variable given, default 8080.
	if port == ":" {
		port = ":8080"
	}

	csv.ProcessFile()

	db.ConnectToDB()
	defer db.DB.Close()

	s := api.CreateNewServer()

	fmt.Println("Server running on port", port)
	http.ListenAndServe(port, s.Router)
}
