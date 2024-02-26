package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello world!")
		d, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Oops", http.StatusBadRequest)
			return

		}

		log.Printf("data %s:", d)
		fmt.Fprintf(w, "hello moji")
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("goodbye world.")
	})

	http.ListenAndServe(":9090", nil)
}
