package handlers

import (
	"net/http"

	"github.com/theluckiestsoul/microservicesgo/product-api/data"
)

func (p *Products) POST(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Hanldle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
	p.l.Printf("[DEBUG] Inserted product: %#v\n", prod)
}
