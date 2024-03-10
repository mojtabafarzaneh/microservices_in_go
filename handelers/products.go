package handelers

import (
	"log"
	"net/http"

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
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProduct(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProduct()
	err := lp.ToJson(w)

	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}

}
