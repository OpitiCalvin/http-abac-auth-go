package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/usecase/user"
	"github.com/OpitiCalvin/http-abac-auth-go/pkg/utils"
	jwt "github.com/dgrijalva/jwt-go"
)

// JwtAuthentication authorizes access to endpoints by analysing jwt token content
func JwtAuthentication(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// list of endpoints that do not require auth
	notAuth := []string{"/api/v1/users", "/api/v1/auth/login"}
	requestPath := r.URL.Path // current request path

	// check if request does not need authorization, serve the request
	for _, value := range notAuth {
		if value == requestPath && r.Method == "POST" {
			next(w, r)
		}
	}

	// response := make(map[string]interface{})

	tokenHeader := r.Header.Get("Authorization") // Grab the token from the header

	if tokenHeader == "" {
		// Token is missing, returns error code 403 Unauthorized
		response := utils.Message(false, "Missing authentication token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		utils.Respond(w, response)
		return
	}

	// The token normally comes in format `Bearer {token-body}`,
	// so we check if the retrieved token matched this requirement
	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		response := utils.Message(false, "Invalid/Malformed authentication token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		utils.Respond(w, response)
		return
	}

	// Grab the token part
	tokenPart := splitted[1]
	tk := &user.Token{}

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKENKEY")), nil
	})

	// Malformed token, returns with http code 403 as usual
	if err != nil {
		response := utils.Message(false, "Malformed authentication token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		utils.Respond(w, response)
		return
	}

	// Token is invalid, maybe not signed on this server
	if !token.Valid {
		response := utils.Message(false, "Token is not valid")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		utils.Respond(w, response)
		return
	}

	// Everythin went well, proceed with request and
	// set the caller to the user retrieved from the parsed token
	_ = fmt.Sprintf("User %v", tk.UserId) // useful for monitoring
	ctx := context.WithValue(r.Context(), "user", tk.UserId)
	r = r.WithContext(ctx)
	next(w, r)
}
