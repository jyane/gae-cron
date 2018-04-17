package main

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/appengine"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}
