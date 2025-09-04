package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/matheusgosk8/goapi/internal/middleware"
)

func Handler (r *chi.Mux){

	//Global Middleware
	r.Use(chimiddle.StripSlashes);
	r.Route("/", func(router chi.Router) {
		router.Get("/", GetPublicRoute)
	})
	r.Route("/account", func(router chi.Router){

		//Middleware for /account route

		router.Use(middleware.Authorization);
		router.Get("/coins", GetCoinBalance)
	})

}