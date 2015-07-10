package main

import (
	"fmt"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Fprintf(w, "<html><head><title>go-test-post</title></head><body><ul>")

	for k, v := range r.PostForm {
		fmt.Printf("%v=%v\n", k, v)
		fmt.Fprintf(w, "<li>%v=%v</li>", k, v)
	}
	fmt.Fprintf(w, "</ul></body></html>")
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
