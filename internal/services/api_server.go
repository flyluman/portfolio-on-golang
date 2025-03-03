package services

import (
	"log"
	"net/http"

	"github.com/flyluman/portfolio-on-golang/internal/handlers"
	"github.com/flyluman/portfolio-on-golang/internal/middleware"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.GetRoot)
	mux.HandleFunc("/whoami", handlers.GetWhoami)
	mux.HandleFunc("/messenger", handlers.PostMessenger)
	mux.HandleFunc("/query", handlers.PostQuery)

	handler := middleware.Logger(mux)

	log.Println("🚀 Server is running on", s.addr)
	return http.ListenAndServe(s.addr, handler)
}
