package health

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type InfoController struct {
}

// Info returns a map containing the health status of the application.
// The map includes a "status" key with a value of "ok".
func (infoController *InfoController) GetInfo(w http.ResponseWriter, r *http.Request) {

	info := map[string]interface{}{
		"status":   "UP",
		"datetime": time.Now().Format(time.RFC3339),
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
func RegisterHealthRoutes(router *mux.Router) {
	controller := InfoController{}
	router.HandleFunc("/pulse/health", controller.GetInfo).Methods("GET")
}
