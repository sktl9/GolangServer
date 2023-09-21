package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world\n"))
}

func main() {
	var PORT = 8081
	var serv = fmt.Sprintf("server start on port %d", PORT)
	http.HandleFunc("/", Handler)

	log.Println(serv)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
}
