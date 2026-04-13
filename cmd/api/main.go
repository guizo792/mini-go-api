package main

import (
	"fmt"
	"net/http"
	
	"github.com/go-chi/chi"
	"github.com/guizo792/mini-go-api/internal/handlers"
	log "github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	loadEnvFile()

	fmt.Println("Starting GO API...")
	fmt.Println(`
 __  __ ___ _   _ ___     ____  ___       _    ____ ___ 
|  \/  |_ _| \ | |_ _|   / ___|/ _ \     / \  |  _ \_ _|
| |\/| || ||  \| || |   | |  _| | | |   / _ \ | |_) | | 
| |  | || || |\  || |   | |_| | |_| |  / ___ \|  __/| | 
|_|  |_|___|_| \_|___|   \____|\___/  /_/   \_\_|  |___|
	`)

	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Error(err)
	}

	fmt.Println("Listening on port 8000...")
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}
