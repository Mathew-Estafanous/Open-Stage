package main

import (
	"log"
	"net/http"
)

const port string = ":8080"

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", HelloWorld)

	log.Printf("Open-Stage starting on port %v", port)
	log.Fatal(http.ListenAndServe(port,r))
}

func HelloWorld(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello World"))
}
