package server

import (
	"fmt"
	"net/http"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/utils"
)

func (s *Server) ApiStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"status": "success",
			"data":   "API running smoothly",
		}
		w.WriteHeader(http.StatusOK)
		// w.Header().Add("Content-Type", "application/json")
		utils.Respond(w, response)
	}
}

func (s *Server) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome")
	}
}
