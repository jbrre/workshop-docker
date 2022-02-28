package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jbrre/workshop-docker/routes"
)

const DEFAULT_PORT = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("PORT not found in env, using default port %s\n", DEFAULT_PORT)
		port = DEFAULT_PORT
	}
	router := mux.NewRouter()
	routes.SetMiddleware(router)
	routes.SetupRoutes(router)

	// err := client.InitDb()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// you'll need to uncomment it for step 4

	s := &http.Server{
		Addr:         ":" + port,
		Handler:      handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router),
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	waiter := make(chan error)
	log.Println("Starting server ...")
	go func() {
		err := s.ListenAndServe()
		waiter <- err
	}()
	log.Printf("AREA server running on http://localhost:%s\n", port)
	err := <-waiter //delete the : once you're at step 4
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server shutted down.")
}
