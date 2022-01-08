package main

import (
	"coffee/app/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()

	// create the handlers
	sm.Handle("/", ph) // products
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:        ":9090",
		Handler:     sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)
	s.ListenAndServe()

	// **Graceful Timeouts** - server will no longer accept requests on Shutdown()
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
