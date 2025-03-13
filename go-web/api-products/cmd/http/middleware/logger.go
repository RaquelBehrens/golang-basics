package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		rw := &customResponseWriter{ResponseWriter: w}

		handler.ServeHTTP(rw, r)

		log.Printf("Method: %s", r.Method)
		log.Printf("URL: %s", r.URL)
		log.Printf("Time: %s", startTime.Format(time.RFC3339))
		log.Printf("Size: %d bytes", rw.size)
	})
}

type customResponseWriter struct {
	http.ResponseWriter
	size int
}

func (rw *customResponseWriter) Write(data []byte) (int, error) {
	bytesWritten, err := rw.ResponseWriter.Write(data)
	rw.size += bytesWritten
	return bytesWritten, err
}
