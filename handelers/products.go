package handelers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/mojtabafarzaneh/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProduct(w, r)
		return
	}
	if r.Method == http.MethodPost {
		p.AddProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {

		p.l.Println("PUT")
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllSubmatch([]byte(r.URL.Path), -1)

		if len(g) != 1 {
			http.Error(w, "bad URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(w, "bad URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(string(idString))
		if err != nil {
			http.Error(w, "bad URI", http.StatusBadRequest)
			return
		}
		p.UpdateProducts(id, w, r)
	}
}
func (p *Products) getProduct(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProduct()
	err := lp.ToJson(w)

	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
		return
	}

}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle post Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "unable to decode json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) UpdateProducts(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "unable to decode json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return

	}

}
