package main

import ("fmt"
		"net/http")


func index_handler(w http.ResponseWriter, r *http.Request){
	// fmt.Fprintf(w, "<h1>Hey There</h1>")
	// fmt.Fprintf(w, "<p>Go is fast!<p>")
	// fmt.Fprintf(w, "<p>... and simple too!<p>")
	// fmt.Fprintf(w, "<p> You %s even and %s<p>", "can", "<strong>variables</strong>")

	fmt.Fprintf(w, "<h1>Hey There</h1> \n <p>Go is fast!<p> \n <p>... and simple too!<p>")
}

func main(){
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
