package handlers

import (
	"microservices-tutorial/data"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
// responses:
//   201: noContent

// DeleteProduct deletes the product from the data store
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle DELETE Products")

	id := p.getID(r)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
