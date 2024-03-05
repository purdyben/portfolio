package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	fmt.Printf(
		"Hello world from %s/%s\n",
		runtime.GOOS,
		runtime.GOARCH,
	)
	//
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
