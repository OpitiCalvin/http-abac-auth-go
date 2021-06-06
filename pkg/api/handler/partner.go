package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/api/presenter"
	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"
	"github.com/OpitiCalvin/http-abac-auth-go/pkg/usecase/partner"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func listPartners(service partner.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errMessage := "Error reading partners"
		var data []*entity.Partner
		var err error

		data, err = service.ListPartners()
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			// w.Write([]byte(errMessage))
			w.Write([]byte(err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errMessage))
			return
		}

		var toJ []*presenter.Partner
		for _, d := range data {
			toJ = append(toJ, &presenter.Partner{
				ID:        d.ID,
				Name:      d.Name,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
			})
		}
		response := map[string]interface{}{
			"status": "success",
			"data":   toJ,
		}
		// if err := json.NewEncoder(w).Encode(toJ); err != nil {
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
	})
}

func createPartner(service partner.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding parner"
		var input struct {
			Name string `json:"name"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		id, err := service.CreatePartner(input.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.Partner{
			ID:   id,
			Name: input.Name,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getPartner(service partner.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading partner"
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := service.GetPartner(id)
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
		toJ := &presenter.Partner{
			ID:        data.ID,
			Name:      data.Name,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		}
		response := map[string]interface{}{
			"status": "success",
			"data":   toJ,
		}
		// if err := json.NewEncoder(w).Encode(toJ); err != nil {
		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func deletePartner(service partner.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error deleting partner"
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		err = service.DeletePartner(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

	})
}

// MakePartnerHandlers mar url handlers
func MakePartnerHandlers(r *mux.Router, n negroni.Negroni, service partner.UseCase) {
	r.Handle("/api/v1/partners", n.With(
		negroni.Wrap(listPartners(service)),
	)).Methods("GET", "OPTIONS").Name("listPartners")

	r.Handle("/api/v1/partners", n.With(
		negroni.Wrap(createPartner(service)),
	)).Methods("POST", "OPTIONS").Name("createPartner")

	r.Handle("/api/v1/partners/{id}", n.With(
		negroni.Wrap(getPartner(service)),
	)).Methods("GET", "OPTIONS").Name("getPartner")

	r.Handle("/api/v1/partners/{id}", n.With(
		negroni.Wrap(deletePartner(service)),
	)).Methods("DELETE", "OPTONS").Name("deletePartner")
}
