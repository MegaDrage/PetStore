package api

import (
	"context"
	"net/http"
	"time"

	"github.com/MegaDrage/PetStore/PetWearablesService/pkg/logger"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(logger *logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(lrw, r)

		logger.Info(
			"HTTP request",
			", method:", r.Method,
			", path: ", r.URL.Path,
			", remote_addr: ", r.RemoteAddr,
			",status: ", lrw.statusCode,
			",duration: ", time.Since(start).String(),
		)
	})
}

type Server struct {
	server *http.Server
	logger *logger.Logger
}

func NewServer(addr string, handler *Handler, logger *logger.Logger) *Server {
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	loggedMux := loggingMiddleware(logger, mux)

	server := &http.Server{
		Addr:         addr,
		Handler:      loggedMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &Server{server: server, logger: logger}
}

func (s *Server) Start() error {
	s.logger.Info("Run HTTP server on: ", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Stopping HTTP server")
	return s.server.Shutdown(ctx)
}