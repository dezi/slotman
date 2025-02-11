package simple

import (
	"net/http"
	"slotman/utils/log"
)

func MuxLog(h http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		logUrl := r.URL.Path
		if len(logUrl) > 48 {
			logUrl = logUrl[:32] + "..." + logUrl[len(logUrl)-16:]
		}

		proto := "http"
		if r.TLS != nil {
			proto = "https"
		}

		log.Printf("Req %s %s %s %s", r.Method, proto, r.RemoteAddr, logUrl)

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
