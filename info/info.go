package info

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type InfoController struct {
}

// Info returns a map containing the health status of the application.
// The map includes a "status" key with a value of "ok".
func (infoController *InfoController) GetInfo(w http.ResponseWriter, r *http.Request) {

	info := map[string]interface{}{
		"info":           os.Getenv("PULSE_MOREINFO"),
		"build":          os.Getenv("BUILD"),
		"commit":         os.Getenv("COMMIT"),
		"branch":         os.Getenv("BRANCH"),
		"buildTimestamp": os.Getenv("BUILD_TIMESTAMP"),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// RegisterDashboardRoutes registers the routes for the health information dashboard.
// It sets up the route for retrieving health information using the provided router.
//
// Parameters:
//   - router (*mux.Router): The router to which the health information route will be added.
//
// Example:
//
//	router := mux.NewRouter()
//	RegisterDashboardRoutes(router)
func RegisterInfoRoutes(router *mux.Router) {
	controller := InfoController{}
	router.HandleFunc("/pulse/info", controller.GetInfo).Methods("GET")
}
