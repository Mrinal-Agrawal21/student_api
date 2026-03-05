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
	"github.com/Mrinal-Agrawal21/student-api/internal/http/handlers/student"
	"github.com/Mrinal-Agrawal21/student-api/internal/storage/sqlite"
)

func main(){
	// load config

	cfg := config.MustLoad()
	
	// database setup
	storage,err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage initialized",slog.String("env",cfg.Env),slog.String("version","1.0.0"))
	// router setup
	router := http.NewServeMux()

	router.HandleFunc("POST /api/v1/students", student.NewStudentHandler(storage))
	router.HandleFunc("GET /api/v1/students/{id}", student.GetStudentByIdHandler(storage))
	// server setup
	server := http.Server{
		Addr: cfg.HTTPServer.Address,
		Handler: router,
	}
	slog.Info("Server is running", slog.String("address", cfg.HTTPServer.Address))

	done := make(chan os.Signal, 1)
	
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func () {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error starting server: %s", err.Error())
		}
	}()

	<-done
	
	slog.Info("Server is shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = server.Shutdown(ctx)	

	if err != nil {
		slog.Error("error shutting down server", slog.String("error", err.Error()))
	}
	slog.Info("Server is shut down")
 
}
