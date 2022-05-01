package server

import (
	"net/http"

	"github.com/cachepump/cachepump/cache"
)

// health is a health check handler.
func health(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(`{"status":"ok"}`))
}

// getCache is a handler for access to cache.
func getCache(rw http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Method is not supported"))
		return
	}

	if key := req.URL.Query().Get("key"); key != "" {
		value, err := cache.Get(key)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(err.Error()))
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write(value)
		return
	}

	rw.WriteHeader(http.StatusBadRequest)
	rw.Write([]byte("Parametor 'key' is not found"))
}
