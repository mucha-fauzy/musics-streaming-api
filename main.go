package main

import (
	"log"
	"musics_streaming_api/internal/repository"
	"musics_streaming_api/internal/services"
	"musics_streaming_api/transport/routes"
	"net/http"
)

func main() {
	// Create a new repository instance
	repo := repository.NewSongImpl()

	// Create a new service instance with the repository
	service := services.ProvideService(repo)

	// Create a new router with the service
	router := routes.NewRouter(service)

	// Set up the routes and start the HTTP server
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router.SetupRoutes()))
}
