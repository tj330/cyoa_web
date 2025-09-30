package routes

import (
	"net/http"

	"github.com/cyoa_web/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	fs := http.FileServer(http.Dir("internal/static/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.Get("/", app.StoryHandler.HanldeGetRoot)
	r.Get("/cyoa/story/{arc}", app.StoryHandler.HandleGetStoryByArc)

	return r
}
