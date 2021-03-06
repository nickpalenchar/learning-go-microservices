package handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h*Hello) ServeHTTP(rw http.ResponseWriter, r*http.Request) {

	h.l.Println("hello world")
	d, err := ioutil.ReadAll(r.Body)

	// error handling should always happen
	if (err != nil) {
		// http pkg has all http status codes as constants
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	log.Printf("data is %s\n", d)
	// writing to http.ResponseWriter automatically sends 
	fmt.Fprintf(rw, "Hello %s", d)
}