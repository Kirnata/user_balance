package apiserver

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

// APIServer ...
type APIServer struct {
	config *Config
	router *mux.Router
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	s.configureRouter()
	log.Println("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}
