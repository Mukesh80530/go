package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1> Whoa, GO is neat!</h1>")
}

func main(){
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":800", nil)
}