package main

import (
	"log"
	"net/http"

	"github.com/DenisGabriel017/fiscalsync-go/internal/platform/config"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não foi encontrado. ")
	}

	cfg := config.Load()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{"status":"ok"}`))
	})

	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, mux))

}
