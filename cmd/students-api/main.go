package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Error while starting the server")
		}
	}()

	//* Until we receive signal keep the server running
	<-done

	//* Graceful shutdown

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("error while shutting down the server", slog.String("error", err.Error()))
	}
	slog.Info("server stopped successfully")

}
