package api

import (
	"encoding/json"
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

	r.Method("GET", "/test", makeHTTPHandlerFunc(s.HandleGetTest))

	return http.ListenAndServe(s.listenAddr, r)
}

func (s *APIServer) HandleGetTest(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, "test successful")
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
