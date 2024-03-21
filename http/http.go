package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pong)
	log.Println("Starting http server ...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Println("pong")
	w.Write([]byte("pong"))
}
