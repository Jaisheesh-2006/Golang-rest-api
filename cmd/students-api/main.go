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
	"github.com/Jaisheesh-2006/Golang-rest-api/internal/http/handlers/student"
	"github.com/Jaisheesh-2006/Golang-rest-api/internal/storage/sqlite"
	"github.com/fatih/color"
)

func main() {
	//* load configuration
	cfg := config.MustLoad()

	//* database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatalf("Error while connecting to the database: %v", err)
	}

	slog.Info("Database initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	//* setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))

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
