package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/theluckiestsoul/microservicesgo/product-api/data"
)

func (p *Products) PUT(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(rw, "Invalid product id", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("Prod: %#v", prod)
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Prodct not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Prodct not found", http.StatusInternalServerError)
		return
	}
	p.l.Printf("[DEBUG] Updated product: %#v\n", prod)
}
