package handlers

import (
	"coffee/app/data"
	"errors"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// handle an add
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// handle an edit

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		p.l.Printf("%v", g)
		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, "Cannot parse id", http.StatusBadRequest)
		}

		p.l.Printf("Got id %v", id)

		p.updateProducts(id, rw, r)

		return
	}

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod, err := productFromBody(r.Body)
	if err != nil {
		http.Error(rw, "cannot parse request", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	prod, err := productFromBody(r.Body)
	if err != nil {
		http.Error(rw, "cannot parse request", http.StatusBadRequest)
	}

	p.l.Printf("PUT Prod: %#v", prod)
	updateErr := data.UpdateProduct(id, prod)

	if updateErr == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}
	if updateErr != nil {
		http.Error(rw, "Something went wring", http.StatusInternalServerError)
	}
}

func productFromBody(body io.ReadCloser) (*data.Product, error) {
	prod := &data.Product{}
	err := prod.FromJSON(body)
	if err != nil {
		return nil, errors.New("cannot create product from JSON")
	}
	return prod, nil
}
