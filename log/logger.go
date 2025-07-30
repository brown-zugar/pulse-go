package log

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger
var atomicLevel zap.AtomicLevel

/*
* init initializes the logger with the default configuration.
 */
func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""

	var err error
	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig
	atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	config.Level = atomicLevel
	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

/*
* Info logs a message at InfoLevel. The message includes any fields passed at the log site, as well as any fields accumulated on the logger.
 */
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

/*
* Debug logs a message at DebugLevel. The message includes any fields passed at the log site, as well as any fields accumulated on the logger.
 */
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

/*
* Error logs a message at ErrorLevel. The message includes any fields passed at the log site, as well as any fields accumulated on the logger.
 */
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

// RegisterLoggerRoutes registers the routes for logger level management.
// It sets up the following endpoints:
// - GET /health/info: Retrieves the current logger level.
// - POST /health/info: Sets the logger level.
func RegisterLoggerRoutes(router *mux.Router) {
	router.HandleFunc("/pulse/logger", getLoggerLevel).Methods("GET")
	router.HandleFunc("/pulse/logger", setLoggerLevel).Methods("POST")
}

// getLoggerLevel handles HTTP requests to retrieve the current logging level.
// It writes the logging level as a JSON response to the provided http.ResponseWriter.
//
// Parameters:
//   - w: http.ResponseWriter to write the JSON response containing the logging level.
//   - r: *http.Request representing the incoming HTTP request.
func getLoggerLevel(w http.ResponseWriter, r *http.Request) {
	level := atomicLevel.Level()
	response := map[string]string{"level": level.String()}
	json.NewEncoder(w).Encode(response)
}

// setLoggerLevel is an HTTP handler that sets the logging level for the application.
// It expects a JSON payload in the request body with a "level" field specifying the desired log level.
// If the "level" field is missing or invalid, it responds with a 400 Bad Request status.
// On success, it sets the log level and responds with a 204 No Content status.
//
// Example request body:
//
//	{
//	  "level": "debug"
//	}
//
// Supported log levels are: "debug", "info", "warn", "error", "dpanic", "panic", and "fatal".
func setLoggerLevel(w http.ResponseWriter, r *http.Request) {
	var request map[string]string
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	level, ok := request["level"]
	if !ok {
		http.Error(w, "Missing level field", http.StatusBadRequest)
		return
	}
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		http.Error(w, "Invalid level", http.StatusBadRequest)
		return
	}
	atomicLevel.SetLevel(zapLevel)
	w.WriteHeader(http.StatusNoContent)
}
