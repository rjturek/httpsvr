package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Printf("%T", r.URL.Path)
	parts := strings.Split(r.URL.Path, "/")
	for _, value := range parts[1:] {
		fmt.Println(value)
	}
}

func main() {
	port := os.Args[1]
	fmt.Println("Listen on port: " + port)
	http.HandleFunc("/", handler)
	portPart := ":" + port
	fmt.Println(portPart)
	http.ListenAndServe(portPart, nil)
}
