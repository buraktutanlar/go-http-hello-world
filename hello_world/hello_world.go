package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	port = ":80"
)

var calls = 0

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	calls++

	time.Sleep(40 * time.Second)

	fmt.Fprintf(w, "Hello, world! You have called me %d times.\n", calls)
}

func init() {
	fmt.Printf("Started server at http://localhost%v.\n", port)
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(port, nil)
}

func main() {}
