package data

import (
	"testing"
)

func TestProductValidation(t *testing.T) {
	p := &Product{
		Name:  "moj",
		Price: 1,
		SKU:   "dofha-jaafh-sdfs",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
