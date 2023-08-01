package routes

import (
	"musics_streaming_api/docs"
	"musics_streaming_api/internal/services"
	"net/http"

	"musics_streaming_api/internal/handlers"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi"
)

type Router struct {
	Handler *handlers.SongHandler
}

func NewRouter(service services.Service) *Router {
	handler := handlers.NewSongHandler(service)
	return &Router{
		Handler: handler,
	}
}

func (r *Router) SetupRoutes() http.Handler {
	mux := chi.NewRouter()
	// Add the Swagger UI route
	docs.SwaggerInfo.Title = "Musics Streaming API"
	docs.SwaggerInfo.Version = "v1"
	conf := httpSwagger.URL("http://localhost:8080/swagger/doc.json")
	mux.Get("/swagger/*", httpSwagger.Handler(conf))

	mux.Route("/api/v1", func(rc chi.Router) {
		rc.Post("/songs", r.Handler.CreateSongHandler)
		rc.Get("/songs", r.Handler.GetSongHandler)
		rc.Put("/songs/{id}", r.Handler.UpdateSongHandler)
		rc.Delete("/songs/{id}", r.Handler.DeleteSongHandler)
	})
	return mux
}
