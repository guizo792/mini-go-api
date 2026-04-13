package main

import (
	"fmt"
	"net/http"
	
	"github.com/go-chi/chi"
	"github.com/guizo792/mini-go-api/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

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
