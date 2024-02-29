package data

import (
	"time"
)

type Product struct {
	ID          int
	name        string
	description string
	price       int
	SKU         string
	createdOn   string
	updatedOn   string
	deletedOn   string
}

var productList = []*Product{
	&Product{
		ID:          1,
		name:        "latte",
		description: "Forthy milky coffee",
		price:       89,
		SKU:         "abc123",
		createdOn:   time.Now().UTC().String(),
		updatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		name:        "Espersso",
		description: "short and strong coffee without milk",
		price:       50,
		SKU:         "dfg456",
		createdOn:   time.Now().UTC().String(),
		deletedOn:   time.Now().UTC().String(),
	},
}
