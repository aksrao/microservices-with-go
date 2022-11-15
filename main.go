package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("error"))
			return
		}
		fmt.Fprintf(rw, "Hello %s", d)
	})
	http.ListenAndServe("127.0.0.1:8081", nil)
}
