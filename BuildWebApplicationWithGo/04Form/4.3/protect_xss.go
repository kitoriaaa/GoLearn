package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("r.Form", r.Form)
	fmt.Println("r.Path", r.URL.Path)
	fmt.Println("r.scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", v)
	}
	fmt.Fprintf(w, "Hello kitoria!")
}

func form(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("form.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println(r.Form.Get("gender"))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
		// fmt.Fprintf(w, r.Form.Get("username"))
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/form", form)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenandServe: ", err)
	}
}
