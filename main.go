package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mojtabafarzaneh/handelers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//create a new serve myx and register the handlers
	ph := handelers.NewProducts(l)
	hh := handelers.NewHello(l)
	gh := handelers.NewGoodbye(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	sm.Handle("/product", ph)

	//create a server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	//start the server
	go func() {
		l.Println("starting server on the port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}

	}()

	//trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	//block until a signal is received
	sig := <-c
	log.Println("got signal:", sig)

	//gracefully sutdown the server, wating max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
