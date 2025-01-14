package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Config struct {
	Message string `json:"message"`
}

var val int
var msg string
var cfg Config

func ReadConfig(quite bool) {
	jsonFile, err := os.Open("config/config.json")
	if err != nil {
		log.Println(err)
		return
	}
	defer jsonFile.Close()

	if !quite {
		log.Println("Successfully Opened config.json")
	}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &cfg)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	ReadConfig(true)
	val++
	if cfg.Message == "" {
		// No Config File so lets status an error
		w.WriteHeader(http.StatusTeapot)
	}
	fmt.Fprintf(w, "serv Test program\n<%v> %v, %v", cfg.Message, msg, val)
}

func main() {
	ReadConfig(false)
	var port int
	flag.StringVar(&msg, "msg", "Hello World", "message to write out")
	flag.IntVar(&port, "port", 10100, "port to listen to")
	flag.Parse()
	log.Printf("listening to port %v\n", port)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
