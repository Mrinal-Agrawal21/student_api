package main

import (
	"context"
	// "fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Mrinal-Agrawal21/student-api/internal/config"
)

func main(){
	// load config

	cfg := config.MustLoad()
	
	// database setup

	// router setup
	router := http.NewServeMux()

	router.HandleFunc("GET /", func (res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Welcome to students api"))
	})
	// server setup
	server := http.Server{
		Addr: cfg.HTTPServer.Address,
		Handler: router,
	}
	slog.Info("Server is running", slog.String("address", cfg.HTTPServer.Address))

	done := make(chan os.Signal, 1)
	
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func () {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("error starting server: %s", err.Error())
		}
	}()

	<-done
	
	slog.Info("Server is shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err := server.Shutdown(ctx)	

	if err != nil {
		slog.Error("error shutting down server", slog.String("error", err.Error()))
	}
	slog.Info("Server is shut down")

}
