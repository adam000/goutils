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
	//muxX.HandleStaticDir(mux, "/test", "./static/deep/folder", handleNotFound, true)
	// Generally, you just want '.' as the second param if you're mirroring
	// the relative path on disk. Note that a param of "/" does not work for
	// the first parameter.
	muxX.HandleStaticDir(mux, "/static", ".", handleNotFound, true)

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
