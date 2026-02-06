package main

import (
	"log"
	"net/http"

	"go-service/internal/auth"
	"go-service/internal/config"
	"go-service/internal/handlers"
	"go-service/internal/students"
)

func main() {
	cfg := config.Load()

	log.Println("Starting Go PDF service...")
	log.Printf("Backend: %s\n", cfg.BackendBaseURL)

	authClient, err := auth.NewClient(cfg.BackendBaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := authClient.Login(cfg.BackendEmail, cfg.BackendPassword); err != nil {
		log.Fatal(err)
	}
	log.Println("Authenticated with backend successfully")
	authClient.DebugCookies()

	studentClient := students.NewClient(
		cfg.BackendBaseURL,
		authClient.HTTP,
		authClient.CSRFToken,
	)

	studentHandler := handlers.NewStudentHandler(studentClient)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/students/", studentHandler.GetByID)

	addr := ":" + cfg.ServicePort
	log.Printf("Listening on %s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
