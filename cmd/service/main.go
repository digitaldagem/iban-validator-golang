package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"iban-validator-golang/src"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/validate_iban", func(r chi.Router) {
		r.Get("/{iban}", func(w http.ResponseWriter, r *http.Request) {
			iban := chi.URLParam(r, "iban")
			response := src.NewIBANValidatorService().ValidateIBAN(iban)
			w.WriteHeader(response.StatusCode)
			_, err := w.Write([]byte(response.Body))
			if err != nil {
				log.Println("error writing the data to the connection as part of an HTTP reply", err)
			}
		})
	})

	log.Println("Service starting on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("error listening to the TCP network address addr and calling Serve", err)
	}
}
