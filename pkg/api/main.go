package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/api/middleware"
	"github.com/OpitiCalvin/http-abac-auth-go/pkg/usecase/partner"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/api/handler"
	"github.com/OpitiCalvin/http-abac-auth-go/pkg/repository"
	"github.com/urfave/negroni"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	partnerRepo := repository.NewPartnerDB(db)
	partnerService := partner.NewPartnerService(partnerRepo)

	r := mux.NewRouter()
	// handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	// partner
	handler.MakePartnerHandlers(r, *n, partnerService)

	http.Handle("/", r)
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"status": "success",
			"data":   "API running smoothly",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		// Addr:         ":" + strconv.Itoa(config.API_PORT),
		Addr: ":8080",
		// Handler:      context.ClearHandler(http.DefaultServeMux),
		Handler:  http.DefaultServeMux,
		ErrorLog: logger,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
