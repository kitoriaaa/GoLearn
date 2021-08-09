package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9090",
	}
	http.HandleFunc("/", body)
	server.ListenAndServe()
}

// have to use post method
// ex) curl -id "name=xxx&mail=yyy" 127.0.0.1:9090/body
