package data

import (
	"time"
	"io"
	"encoding/json"
)



type Product struct {
	ID int `json:"id"` // see **struct tags** in notes
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string `json:"sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

type Products []*Product

// uses Encode from json to encode json on an io.Writer
// the io.Writer is usually the http.ResponseWriter (which
// implements io.Writer)
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = Products {
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Foamy milk coffee",
		Price: 2.45,
		SKU: "abc123",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 1,
		Name: "Espresso",
		Description: "Short and strong coffee without milk",
		Price: 2.45,
		SKU: "abc123",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}