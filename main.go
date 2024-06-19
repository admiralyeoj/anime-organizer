package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/admiralyeoj/anime-organizer/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	dbg := flag.Bool("debug", false, "Enable debug mode")

	flag.Parse()

	if *dbg {
		fmt.Println("Debug mode is enabled")
		err := os.Remove(dbPath)
		if err != nil {
			fmt.Println("Error deleting the file:", err)
			return
		}
	}

	const filepathRoot = "./app"
	const port = "8080"

	// apiCfg := apiConfig{}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fsHandler := http.StripPrefix("/", http.FileServer(http.Dir(filepathRoot)))

	r.Handle("/", fsHandler)
	r.Handle("/*", fsHandler)

	// Mount API Routes Here
	apiRouter := chi.NewRouter()
	r.Mount("/api", apiRouter)

	apiRouter.Get("/get-series", controllers.GetSeriesHandler)

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	http.ListenAndServe(":"+port, r)
}
