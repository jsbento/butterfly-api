package server

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	HTTPServer  *http.Server
	MongoClient *mongo.Client
}

func (s *Server) Start() error {
	log.Printf("Server started on port %s", s.HTTPServer.Addr)
	return s.HTTPServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.MongoClient.Disconnect(ctx); err != nil {
		return err
	}
	if err := s.HTTPServer.Close(); err != nil {
		return err
	}
	log.Printf("Server stopped")
	return nil
}
