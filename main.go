package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", oofServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func oofServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello oof!")
}
