package service

import (
	"github.com/brown-zugar/pulse-go/health"
	"github.com/brown-zugar/pulse-go/info"
	"github.com/brown-zugar/pulse-go/log"
	"github.com/gorilla/mux"
)

var router *mux.Router

// init initializes the application by setting up the router and applying middleware.
// It logs an informational message indicating that the application is starting.
func loadRoutes() {
	log.Info("Createting API /pulse")
	router = mux.NewRouter()

	// ************** API /pulse **************
	health.RegisterHealthRoutes(router)
	log.RegisterLoggerRoutes(router)
	info.RegisterInfoRoutes(router)
}

// Enable initializes the Pulse application, loads all routes, and returns the configured HTTP router.
func Enable() *mux.Router {
	log.Info("Pulse Application starting...")
	loadRoutes()
	return router
}
