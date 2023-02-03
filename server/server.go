package server

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/jsbento/butterfly-api/mongo"
)

func NewServer(ctx context.Context, dbUri string) *Server {
	r := mux.NewRouter()

	s := &Server{
		HTTPServer: &http.Server{
			Addr:    ":8080",
			Handler: r,
		},
		MongoClient: m.NewMongoClient(ctx, dbUri),
	}

	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods(http.MethodGet)

	r.HandleFunc("/shipments", s.CreateShipment()).Methods(http.MethodPost)
	r.HandleFunc("/shipments/{id}", s.UpdateShipment()).Methods(http.MethodPut)
	r.HandleFunc("/shipments/{id}", s.GetShipment()).Methods(http.MethodGet)
	r.HandleFunc("/shipments", s.SearchShipments()).Methods(http.MethodGet)

	r.HandleFunc("/organizations", s.CreateOrganization()).Methods(http.MethodPost)
	r.HandleFunc("/organizations/{id}", s.UpdateOrganization()).Methods(http.MethodPut)
	r.HandleFunc("/organizations/{id}", s.GetOrganization()).Methods(http.MethodGet)
	r.HandleFunc("/organizations", s.SearchOrganizations()).Methods(http.MethodGet)

	r.HandleFunc("/releases", s.CreateRelease()).Methods(http.MethodPost)
	r.HandleFunc("/releases/{id}", s.UpdateRelease()).Methods(http.MethodPut)
	r.HandleFunc("/releases/{id}", s.GetRelease()).Methods(http.MethodGet)
	r.HandleFunc("/releases", s.SearchReleases()).Methods(http.MethodGet)

	r.Use(mux.CORSMethodMiddleware(r))

	s.HTTPServer.Handler = r

	return s
}
