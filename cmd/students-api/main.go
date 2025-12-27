package main

import (
	"log"
	"net/http"

	"github.com/Jaisheesh-2006/Golang-rest-api/internal/config"
	"github.com/fatih/color"
)

func main() {
	//* load configuration

	cfg := config.MustLoad()

	//* database setup

	//* setup router

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})

	//* start server

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	color.Green("Server started at %s", cfg.Address)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting the server")
	}

}
