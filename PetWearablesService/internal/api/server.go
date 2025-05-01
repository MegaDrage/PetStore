package api

import (
	"context"
	"net/http"
	"time"

	"github.com/MegaDrage/PetStore/PetWearablesService/pkg/logger"
	"github.com/gorilla/mux"
)

type Server struct {
	server *http.Server
	logger *logger.Logger
}

func NewServer(addr string, handler *Handler, logger *logger.Logger) *Server {
	router := mux.NewRouter()
	router.HandleFunc("/pets/{pet_id}/metrics", handler.GetPetMetrics).Methods("GET")

	server := &http.Server{
		Addr:         addr,
		Handler:      router,
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