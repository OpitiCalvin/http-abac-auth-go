package server

import (
	"log"
	"net/http"

	// "novelsTradeIn/pkg/api"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	// userService api.UserService
}

// func NewServer(router *mux.Router, userService api.UserService) *Server {
func NewServer(router *mux.Router) *Server {
	return &Server{
		Router: router,
		// userService: userService,
	}
}

// Run function initializes the application's routes
func (s *Server) Run() error {
	s.InitializeRoutes()

	// run the server through the router
	err := http.ListenAndServe(":8080", s.Router)
	if err != nil {
		log.Printf("There was an error initializing server listening and serving: %v\n", err)
		return err
	}
	log.Printf("Server listening on Port 8080")

	return nil
}
