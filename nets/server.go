package nets

import (
	"net/http"
	"time"
)

// ListenAndServe 并设置Timeout.
func ListenAndServe(addr string, handler http.Handler) error {
	// nolint: gomnd
	server := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       90 * time.Second,
	}

	return server.ListenAndServe()
}
