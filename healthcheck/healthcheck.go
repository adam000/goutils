// Package healthcheck aims to provide simple, consistent healthchecking to
// compatible services.
package healthcheck

import (
	"encoding/json"
	"log"
	"net/http"
)

func NewHealthcheckHandler(healthPoints func() map[string]interface{}) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(healthPoints())
	}
}

func PingHealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(struct {
		IsAlive string
	}{
		"true",
	})
	if err != nil {
		log.Printf("error issuing healthcheck: %v", err)
	}
}
