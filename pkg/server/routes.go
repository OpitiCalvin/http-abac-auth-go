package server

func (s *Server) InitializeRoutes() {
	router := s.Router

	router.HandleFunc("/status", s.ApiStatus()).Methods("GET")
	router.HandleFunc("/", s.IndexHandler()).Methods("GET")
}
