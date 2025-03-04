package services

import (
	"log"
	"net/http"

	"github.com/flyluman/portfolio-on-golang/internal/handlers"
	"github.com/flyluman/portfolio-on-golang/internal/middleware"
)

type APIServer struct {
	addr     string
	certFile string
	keyFile  string
}

func NewAPIServer(addr, certFile, keyFile string) *APIServer {
	return &APIServer{addr: addr, certFile: certFile, keyFile: keyFile}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.GetRoot)
	mux.HandleFunc("GET /whoami", handlers.GetWhoami)
	mux.HandleFunc("POST /messenger", handlers.PostMessenger)
	mux.HandleFunc("POST /query", handlers.PostQuery)

	handler := middleware.Logger(mux)

	log.Println("🚀 Server is running on https://luman.mooo.com", s.addr)
	return http.ListenAndServeTLS(s.addr, s.certFile, s.keyFile, handler)
}
