package pulse

import (
	"github.com/chava-gnolasco/pulse/health"
	"github.com/chava-gnolasco/pulse/info"
	"github.com/chava-gnolasco/pulse/log"
	"github.com/gorilla/mux"
)

// Config holds the configuration for the Pulse module
type Config struct {
	// BasePath is the base path for all pulse endpoints (default: "/pulse")
	BasePath string
	// EnableHealth enables the health endpoint (default: true)
	EnableHealth bool
	// EnableInfo enables the info endpoint (default: true)
	EnableInfo bool
	// EnableLogger enables the logger configuration endpoint (default: true)
	EnableLogger bool
}

// DefaultConfig returns a default configuration for Pulse
func DefaultConfig() *Config {
	return &Config{
		BasePath:     "/pulse",
		EnableHealth: true,
		EnableInfo:   true,
		EnableLogger: true,
	}
}

// Pulse represents the pulse module instance
type Pulse struct {
	config *Config
}

// New creates a new Pulse instance with the given configuration
func New(config *Config) *Pulse {
	if config == nil {
		config = DefaultConfig()
	}
	
	log.Info("Pulse module initializing...")
	
	return &Pulse{
		config: config,
	}
}

// RegisterRoutes registers all enabled pulse routes to the provided router
func (p *Pulse) RegisterRoutes(router *mux.Router) {
	log.Info("Registering Pulse routes...")
	
	if p.config.EnableHealth {
		health.RegisterHealthRoutes(router)
		log.Info("Health routes registered")
	}
	
	if p.config.EnableLogger {
		log.RegisterLoggerRoutes(router)
		log.Info("Logger routes registered")
	}
	
	if p.config.EnableInfo {
		info.RegisterInfoRoutes(router)
		log.Info("Info routes registered")
	}
	
	log.Info("Pulse module routes registration completed")
}

// RegisterRoutesWithRouter creates a new subrouter with the configured base path
// and registers all enabled pulse routes to it
func (p *Pulse) RegisterRoutesWithRouter(router *mux.Router) *mux.Router {
	subrouter := router.PathPrefix(p.config.BasePath).Subrouter()
	
	if p.config.EnableHealth {
		health.RegisterHealthRoutesWithBasePath(subrouter, "")
		log.Info("Health routes registered with base path")
	}
	
	if p.config.EnableLogger {
		log.RegisterLoggerRoutesWithBasePath(subrouter, "")
		log.Info("Logger routes registered with base path")
	}
	
	if p.config.EnableInfo {
		info.RegisterInfoRoutesWithBasePath(subrouter, "")
		log.Info("Info routes registered with base path")
	}
	
	log.Info("Pulse module routes registration with base path completed")
	return subrouter
}
