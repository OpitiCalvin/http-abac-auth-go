package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/api/presenter"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
	"github.com/OpitiCalvin/http-abac-auth-go/pkg/usecase/client"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func listClients(service client.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading clients"
		var data []*entity.Client
		var err error
		var response map[string]interface{}

		w.Header().Set("Content-Type", "application/json")

		data, err = service.ListClients()
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		var toJ []*presenter.Client

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			response = map[string]interface{}{
				"status": "No client records found",
				"data":   toJ,
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		for _, d := range data {
			toJ = append(toJ, &presenter.Client{
				ID:         int64(d.ID),
				ClientName: d.ClientName,
				Products:   d.Products,
				PartnerID:  int64(d.PartnerID),
				CreatedAt:  d.CreatedAt,
				UpdatedAt:  d.UpdatedAt,
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

func createClient(service client.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding a client"
		var input struct {
			ClientName string `json:"client_name"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// w.Write([]byte(errorMessage))
			w.Write([]byte(err.Error()))
			return
		}

		err = service.CreateClient(input.ClientName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// w.Write([]byte(errorMessage))
			w.Write([]byte(err.Error()))
			return
		}

		toJ := &presenter.ClientCreated{
			ClientName: input.ClientName,
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

func updateClient(service client.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating a client"
		var input struct {
			ClientName string `json:"client_name"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// w.Write([]byte(errorMessage))
			w.Write([]byte(err.Error()))
			return
		}

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := service.GetClient(int64(id))
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

		if data.ClientName == input.ClientName {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Update cancelled as new and old client names are similar"))
			return
		}

		data.ClientName = input.ClientName
		err = service.UpdateClient(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// w.Write([]byte(errorMessage))
			w.Write([]byte(err.Error()))
			return
		}

		toJ := &presenter.ClientCreated{
			ClientName: input.ClientName,
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

func getClient(service client.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading client"
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := service.GetClient(int64(id))
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

		toJ := &presenter.Client{
			ID:         data.ID,
			ClientName: data.ClientName,
			Products:   data.Products,
			PartnerID:  data.PartnerID,
			CreatedAt:  data.CreatedAt,
			UpdatedAt:  data.UpdatedAt,
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

func deleteClient(service client.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error deleting client"
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		err = service.DeleteClient(int64(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakeClientHandlers make url handlers for client resources
func MakeClientHandlers(r *mux.Router, n negroni.Negroni, service client.UseCase) {
	r.Handle("/api/v1/clients", n.With(
		negroni.Wrap(listClients(service)),
	)).Methods("GET", "OPTIONS").Name("listClients")

	r.Handle("/api/v1/clients", n.With(
		negroni.Wrap(createClient(service)),
	)).Methods("POST", "OPTIONS").Name("createClient")

	r.Handle("/api/v1/clients/{id}", n.With(
		negroni.Wrap(getClient(service)),
	)).Methods("GET", "OPTIONS").Name("getClient")

	r.Handle("/api/v1/clients/{id}", n.With(
		negroni.Wrap(updateClient(service)),
	)).Methods("PUT", "OPTIONS").Name("updateClient")

	r.Handle("/api/v1/clients/{id}", n.With(
		negroni.Wrap(deleteClient(service)),
	)).Methods("DELETE", "OPTONS").Name("deleteClient")
}
