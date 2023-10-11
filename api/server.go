package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	listenAddr string
}

func NewServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() error {
	r := chi.NewRouter()

	return http.ListenAndServe(s.listenAddr, r)
}
