// Package handlers Product API
//
// Documentation for Product API
//
//    Schemes: http
//    BasePath: /
//    Version: 0.0.1
//
//    Consumes:
//    - application/json
//
//    Produces:
//    - application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"microservices-tutorial/data"
	"net/http"
	"strconv"
)

// A list of products returned to the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []*data.Product
}

// swagger:response noContent
type productsNoContent struct {
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The ID of the product to delete from the data source
	// in: path
	// required: true
	ID int `json:"id"`
}

type Products struct {
	l *log.Logger
}

type KeyProduct struct {
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product")
			http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		if validationErr := prod.Validate(); validationErr != nil {
			p.l.Println("[ERROR] validating product")
			http.Error(
				rw,
				fmt.Sprintf("Error validaing product: %s", validationErr),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}

func (p *Products) getID(r *http.Request) int {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	return id
}
