package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/go-chi/cors"
)

func main() {
	// get the environment variables
	godotenv.Load(".env")

	// check if we get the port from the env variables
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Println("Port:", portString)

	// create a new router
	router := chi.NewRouter()

	// permissions
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	//
	v1Router := chi.NewRouter() 
	v1Router.HandleFunc("/healthz", handlerReadiness)

	router.Mount("/v1", v1Router)

	// create a server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	// check if we ever get an error
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}


}
