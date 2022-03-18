package logging

import (
	"log"
	"net/http"
	"time"

	zerolog "github.com/rs/zerolog/log"
)

func LogMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Got request: method: %s, URL: %s, Elapsed time: %s", r.Method, r.RequestURI, time.Since(start))
		next.ServeHTTP(w, r)
	})
}

func EasyLogInfo(prefix, message, data string) {
	zerolog.Info().
		Str("service", prefix).
		Msgf(message + data)
}

func EasyLogError(prefix, message, data string, reportedErr error) {
	zerolog.Error().
		Err(reportedErr).
		Str("service", prefix).
		Msgf(message + data)
}

func EasyLogFatal(prefix, message, data string, reportedErr error) {
	zerolog.Fatal().
		Err(reportedErr).
		Str("service", prefix).
		Msgf(message + data)
}
