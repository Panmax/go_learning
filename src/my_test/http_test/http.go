package main

import (
	"net/http"
)

func foo(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("foo("))
		next(w, r)
		w.Write([]byte(")"))
	}
}

func bar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bar("))
		next(w, r)
		w.Write([]byte(")"))
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func main() {
	http.Handle("/", foo(bar(test)))
	http.ListenAndServe(":8081", nil)
}
