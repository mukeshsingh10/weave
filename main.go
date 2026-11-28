package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mukeshsingh10/weave/internal/config"
)

type Response struct {
	Status string `json:"status"`
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := &Response{
		Status: "ok",
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to encode health response %v", err)
		return
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using environment variables")
	}
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handleHealth)

	server := &http.Server{
		Addr:    cfg.HTTPPort,
		Handler: mux,
	}
	log.Printf("server is listening on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
