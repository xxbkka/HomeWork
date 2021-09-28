package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	http.HandleFunc("/healthz", healthzHandler)
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Add(k, v[0])
	}
	w.Header().Add("VERSION", os.Getenv("VERSION"))
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(host)
	log.Println("http 返回码")
	w.WriteHeader(200)
}
