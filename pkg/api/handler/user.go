package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/api/presenter"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
	"github.com/OpitiCalvin/http-abac-auth-go/pkg/usecase/user"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func listUsers(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading users"
		var data []*entity.User
		var err error
		var response map[string]interface{}

		w.Header().Set("Content-Type", "application/json")

		data, err = service.ListUsers()
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		var toJ []*presenter.User

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			response = map[string]interface{}{
				"status": "No user records found",
				"data":   toJ,
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		for _, d := range data {
			toJ = append(toJ, &presenter.User{
				ID:        int64(d.ID),
				Email:     d.Email,
				Username:  d.Username,
				ClientID:  d.ClientID,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
			})
		}

		response = map[string]interface{}{
			"status": "success",
			"data":   toJ,
		}
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func createUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding a user"
		var input struct {
			Email    string `json:"email"`
			Username string `json:"username"`
			Password string `json:"password"`
			ClientID int64  `json:"client_id"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		id, err := service.CreateUser(input.Email, input.Username, input.Password, input.ClientID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.User{
			ID:       id,
			Email:    input.Email,
			Username: input.Email,
			ClientID: input.ClientID,
		}

		response := map[string]interface{}{
			"status": "success",
			"data":   toJ,
		}

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading user"
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := service.GetUser(int64(id))
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.User{
			ID:        data.ID,
			Email:     data.Email,
			Username:  data.Username,
			ClientID:  data.ClientID,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		}

		response := map[string]interface{}{
			"status": "success",
			"data":   toJ,
		}

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func deleteUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error deleting user"
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		err = service.DeleteUser(int64(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakeUserHandlers make url handlers for user resources
func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service user.UseCase) {
	r.Handle("/api/v1/clients", n.With(
		negroni.Wrap(listUsers(service)),
	)).Methods("GET", "OPTIONS").Name("listUsers")

	r.Handle("/api/v1/users", n.With(
		negroni.Wrap(createUser(service)),
	)).Methods("POST", "OPTIONS").Name("createUser")

	r.Handle("/api/v1/users/{id}", n.With(
		negroni.Wrap(getUser(service)),
	)).Methods("GET", "OPTIONS").Name("getUser")

	r.Handle("/api/v1/users/{id}", n.With(
		negroni.Wrap(deleteUser(service)),
	)).Methods("DELETE", "OPTONS").Name("deleteUser")
}
