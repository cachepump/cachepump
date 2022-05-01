package server

import (
	"context"
	"net/http"
	"time"

	logger "github.com/AntonYurchenko/log-go"
)

// Routing is a entity for definition of mapping URL to handler.
type Routing map[string]func(http.ResponseWriter, *http.Request)

func (r *Routing) init() {
	if r != nil {
		for path, handler := range *r {
			http.HandleFunc(path, handler)
			logger.DebugF("Initialisation handler for path: %q", path)
		}
	}
}

// srv is a custom http server.
var srv *http.Server = nil

// Start ups a http server on selected endpoint.
func Start(endpoint string) {

	srv = &http.Server{Addr: endpoint}
	routing.init()

	logger.InfoF("Endpoint of http server is %q", endpoint)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.FatalF("Http server has been stopped with error: %v", err)
		}
		logger.Info("Http server has been stopped. Have a god day!")
	}()
}

// Stop down a http server.
func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if srv != nil {
		logger.Info("Stopping a http server ...")
		srv.Shutdown(ctx)
	}
}
