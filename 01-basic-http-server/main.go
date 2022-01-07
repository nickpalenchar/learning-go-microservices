package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
		log.Println("hello world")
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
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye world")
	})

	http.ListenAndServe(":9090", nil)
}
