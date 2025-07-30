package main

import (
	"net/http"

	"github.com/chava-gnolasco/pulse/health"
	"github.com/chava-gnolasco/pulse/info"
	"github.com/chava-gnolasco/pulse/log"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var router *mux.Router

// init initializes the application by setting up the router and applying middleware.
// It logs an informational message indicating that the application is starting.
func init() {
	log.Info("Pulse Application starting...")
	router = mux.NewRouter()

	// ************** API /pulse **************
	health.RegisterHealthRoutes(router)
	log.RegisterLoggerRoutes(router)
	info.RegisterInfoRoutes(router)
}

// main starts the HTTP server on port 8080 using the provided router.
// If the server fails to start, it logs the error using zap.
func main() {

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Error("Server failed to start: ", zap.Error(err))
	}
}
