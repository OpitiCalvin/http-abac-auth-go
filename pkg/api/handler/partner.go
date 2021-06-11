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
		errorMessage := "Error reading partners"
		var data []*entity.Partner
		var err error
		var response map[string]interface{}

		w.Header().Set("Content-Type", "application/json")

		data, err = service.ListPartners()
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			// w.Write([]byte(err.Error()))
			return
		}

		var toJ []*presenter.Partner

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			// w.Write([]byte(errorMessage))
			response = map[string]interface{}{
				"status": "No Partner records found",
				"data":   toJ,
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		for _, d := range data {
			toJ = append(toJ, &presenter.Partner{
				ID:          int64(d.ID),
				PartnerName: d.PartnerName,
				CreatedAt:   d.CreatedAt,
				UpdatedAt:   d.UpdatedAt,
			})
		}
		response = map[string]interface{}{
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

func createPartner(service partner.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding parner"
		var input struct {
			PartnerName string `json:"partner_name"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			// log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		id, err := service.CreatePartner(input.PartnerName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.PartnerCreated{
			ID:          int64(id),
			PartnerName: input.PartnerName,
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

		data, err := service.GetPartner(int64(id))
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
			ID:          int64(data.ID),
			PartnerName: data.PartnerName,
			CreatedAt:   data.CreatedAt,
			UpdatedAt:   data.UpdatedAt,
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

		err = service.DeletePartner(int64(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

	})
}

// MakePartnerHandlers make url handlers for partner resources
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
