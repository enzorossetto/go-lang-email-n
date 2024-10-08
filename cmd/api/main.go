package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	db := database.NewDb()
	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{Db: db},
	}
	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)

		r.Post("/", endpoints.ErrorHandler(handler.CampaignPost))
		r.Get("/", endpoints.ErrorHandler(handler.CampaignGet))
		r.Get("/{id}", endpoints.ErrorHandler(handler.CampaignGetById))
		r.Delete("/{id}", endpoints.ErrorHandler(handler.CampaignDelete))
	})

	http.ListenAndServe(":3000", r)
}
