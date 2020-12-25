package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
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
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("form.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// tokenの合法性を確かめる
		} else {
			// tokenが存在しなければエラーを出す
		}
		fmt.Println("usrname length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
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
