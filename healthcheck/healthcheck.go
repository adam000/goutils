package healthcheck

import (
	"encoding/json"
	"net/http"
)

// The goal of this package is to provide simple, consistent health checking to
// my services.

func NewHealthcheckHandler(healthPoints func() map[string]interface{}) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(healthPoints())
	}
}

func PingHealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		isAlive string
	}{
		"true",
	})
}
