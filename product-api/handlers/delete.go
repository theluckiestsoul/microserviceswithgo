package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/theluckiestsoul/microservicesgo/product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
// responses:
// 	200: noContent

// DeleteProducts deletes a product from database
func (p *Products) DeleteProducts(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(rw, "Invalid product id", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle DELETE Product")
	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Prodct not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Prodct not found", http.StatusInternalServerError)
		return
	}
	p.l.Printf("[DEBUG] Deleted product with id: %#v\n", id)
}
