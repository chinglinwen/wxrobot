package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi there, this is from backend")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	fmt.Println("body:", string(body))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
