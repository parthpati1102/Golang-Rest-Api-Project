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

	"github.com/parthpati1102/Golang-pRest-API-Project/internal/config"
	"github.com/parthpati1102/Golang-pRest-API-Project/internal/http/handlers/student"
	sqlite "github.com/parthpati1102/Golang-pRest-API-Project/internal/storage/sqllite"
)

func main() {
	//load Config
	cfg := config.MustLoad()

	//database setup
	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage Initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))

	//setup Server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server Started", slog.String("Address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start Server")
		}
	}()

	<-done

	slog.Info("Shutting down the Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("Failed to shutdown Server", slog.String("error", err.Error()))
	}

	slog.Info("Server Shutdown Successfully")

}
