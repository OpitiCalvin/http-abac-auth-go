package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/utils"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// List of endpoints that do not require auth
		notAuth := []string{"/api/user/new", "/api/user/login"}
		requestPath := r.URL.Path // current request path

		// check if request does not need authentication, serve the request if it doesn't
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") // Grab the token from the header

		if tokenHeader == "" {
			// Token is missing, returns error code 403 Unauthorized
			response = utils.Message(false, "Missing authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		// The token normally comes in format `Bearer {token-body}`,
		// so we check if the retrieved token matched this requirement
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = utils.Message(false, "Invalid/Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		// Grab the token part, what we are truly interested in
		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKENKEY")), nil
		})

		// Malformed token, returns with http code 403 as usual
		if err != nil {
			response = utils.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		// Token is invalid, maybe not signed on this server
		if !token.Valid {
			response = utils.Message(false, "Token is not valid")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		// Everything went well, proceed with request and set the caller to the user retrieved from the parsed token
		fmt.Sprintf("User %", tk.Username) // Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) // Proceed in the middleware chain!

	})
}
