package main

import (
	"fmt"
	"net/http"
)


type Str string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}


func (s Str) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	response := fmt.Sprintf("%v", s)
	fmt.Fprint(w, response)
}



func (x Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	response := fmt.Sprintf("%v%v%v", x.Greeting,x.Punct,x.Who)
	fmt.Fprint(w, response)
}


func main() {
	go http.Handle("/", Str("This is sparta!!!"))
	http.Handle("/string", Str("I'm a frayed knot."))
	http.Handle("/struct", Struct{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}