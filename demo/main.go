package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("received a request")
	fmt.Fprintf(w, "Hello, I am an ingress-controller demo!\n")
}

func main() {
	s := &http.Server{
		Addr:      ":8080",
		Handler:   &myhandler{},
		TLSConfig: &tls.Config{
			// ServerName: "www.simple.org",
		},
	}

	err := s.ListenAndServeTLS("../certs/simple.org.crt", "../certs/simple.org.key")
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
	fmt.Println("ListenAndServeTLS exit ok")
}
