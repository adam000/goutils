package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adam000/goutils/muxX"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	// This one only lets you look at the file in the test subfolder
	muxX.HandleStaticDir(mux, "/test", "./static/deep/folder", handleNotFound, true)
	// Generally, you just want '.' as the second param if you're mirroring
	// the relative path on disk. Note that a param of "/" does not work for
	// the first parameter. In this example, it exposes "main.go" since that's
	// in the current working directory if you `go run main.go`.
	muxX.HandleStaticDir(mux, "/root", ".", handleNotFound, true)
	// Another example: you want to handle a nearby directory (`./static`) as
	// if it's at the root of your mux (`/static`)
	// This allows you to allow certain directories at the root of your project
	// without exposing the executable.
	muxX.HandleStaticDir(mux, "/static", "./static", handleNotFound, true)

	log.Println("Starting...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	log.Println("Tried to load a dir and we don't want that")
	fmt.Fprintf(w, "Not found! %s", r.URL.Path)
}
