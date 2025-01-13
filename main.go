package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
)

var val int
var msg  string

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v, %v\nHello from version 3", msg, val)
	val++
}

func main() {
	flag.StringVar(&msg, "msg", "Hello World", "message to write out")
	flag.Parse()
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10100", nil))
}
