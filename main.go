package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Message string `json:"message"`
}

var val int
var msg string
var cfg Config

func ReadConfig() {
	jsonFile, err := os.Open("config/config.json")
	if err != nil {
		log.Println(err)
		return
	}
	defer jsonFile.Close()

	log.Println("Successfully Opened config.json")
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &cfg)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	ReadConfig()
	val++
	if cfg.Message == "" {
		// No Config File so lets status an error
		w.WriteHeader(http.StatusTeapot)
	}
	fmt.Fprintf(w, "serv Test program\n<%v> %v, %v", cfg.Message, msg, val)
}

func main() {
	ReadConfig()
	flag.StringVar(&msg, "msg", "Hello World", "message to write out")
	flag.Parse()
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10100", nil))
}
