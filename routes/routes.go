package routes

import (
	"github.com/agile-work/srv-mdl-task/controllers"
	"github.com/go-chi/chi"
)

// Setup configure the API endpoints
func Setup() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/task", func(r chi.Router) {
		r.Get("/admin/features", controllers.Features)
	})

	return router
}
