// Ping provides a stupid simple healthchecker, with an optional argument for port.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/adam000/goutils/healthcheck"
	"github.com/gorilla/mux"
)

func main() {
	port := ":8080"
	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Could not convert %s to string:%v", os.Args[1], err)
		}
		if i <= 1024 || i > 65535 {
			log.Fatalf("Port %d outside of valid range (1024, 65535]", i)
		}
		port = fmt.Sprintf(":%d", i)
	}
	r := mux.NewRouter()

	r.HandleFunc("/health", healthcheck.PingHealthcheckHandler)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(port, nil))
}
