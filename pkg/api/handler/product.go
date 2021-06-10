package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/api/presenter"
	"github.com/OpitiCalvin/http-abac-auth-go/pkg/usecase/product"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func listProducts(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errMessage := "Error reading products"
		var data []*entity.Product
		var err error
		var response map[string]interface{}

		w.Header().Set("Content-Type", "application/json")

		data, err = service.ListProducts()
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
		var toJ []*presenter.Product

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			// w.Write([]byte(errMessage))
			response = map[string]interface{}{
				"status": "No Product records found",
				"data":   toJ,
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		for _, d := range data {
			toJ = append(toJ, &presenter.Product{
				ID:        int64(d.ID),
				Name:      d.Name,
				BaseURL:   d.BaseURL,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
			})
		}
		response = map[string]interface{}{
			"status": "success",
			"data":   toJ,
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			// log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
	})
}

func createProduct(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding product"
		var input struct {
			Name    string `json:"name"`
			BaseURL string `json:"base_url"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		id, err := service.CreateProduct(input.Name, input.BaseURL)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.Product{
			ID:      int64(id),
			Name:    input.Name,
			BaseURL: input.BaseURL,
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

func getProduct(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading product"
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := service.GetProduct(int64(id))
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
		toJ := &presenter.Product{
			ID:        int64(data.ID),
			Name:      data.Name,
			BaseURL:   data.BaseURL,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		}
		response := map[string]interface{}{
			"status": "success",
			"data":   toJ,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func deleteProduct(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error deleting product"
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		err = service.DeleteProduct(int64(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

	})
}

// MakeProductHandlers make url handlers
func MakeProductHandlers(r *mux.Router, n negroni.Negroni, service product.UseCase) {
	r.Handle("/api/v1/products", n.With(
		negroni.Wrap(listProducts(service)),
	)).Methods("GET", "OPTIONS").Name("listProducts")

	r.Handle("/api/v1/products", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS").Name("createProduct")

	r.Handle("/api/v1/products/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS").Name("getProduct")

	r.Handle("/api/v1/products/{id}", n.With(
		negroni.Wrap(deleteProduct(service)),
	)).Methods("DELETE", "OPTONS").Name("deleteProduct")
}
