package proxy

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"slotman/utils/log"
	"time"
)

func (sv *Service) startServers() (err error) {

	if sv.httpRunning {
		return
	}

	err = sv.startHTTP()
	log.Cerror(err)

	return
}

func (sv *Service) stopServers() (err error) {

	if !sv.httpRunning {
		return
	}

	err = sv.stopHTTP()
	log.Cerror(err)

	return
}

func (sv *Service) startHTTP() (err error) {

	if sv.httpRunning {
		return
	}

	sv.httpMux = &http.ServeMux{}
	sv.httpMux.HandleFunc("/", sv.handleIndex)
	sv.httpMux.HandleFunc("/ws", sv.handleWs)

	sv.httpServer = &http.Server{
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       3600 * time.Second,
		WriteTimeout:      3600 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	sv.httpServer.Addr = ":8877"
	sv.httpServer.Handler = sv.httpMux

	go func() {

		log.Printf("Starting HTTP server on %s", sv.httpServer.Addr)
		sv.httpRunning = true

		err := sv.httpServer.ListenAndServe()
		if err != nil && err.Error() != "http: Server closed" {
			log.Cerror(err)
		}

		log.Printf("Stopped HTTP server on %s", sv.httpServer.Addr)
		sv.httpRunning = false
	}()

	return
}

func (sv *Service) stopHTTP() (err error) {

	if sv.httpServer == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	err = sv.httpServer.Shutdown(ctx)
	if err != nil {
		log.Cerror(err)
		return
	}

	sv.httpServer = nil
	sv.httpRunning = false

	return
}

func (sv *Service) handleIndex(w http.ResponseWriter, r *http.Request) {

	_ = r

	helloPage := `<html><body><center><h1>Welcome to Proxy!</h1><h3>(%s)</h3></center></body></html>`
	output := fmt.Sprintf(helloPage, time.Now())
	_, _ = io.WriteString(w, output)
}
