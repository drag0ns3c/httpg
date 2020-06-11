package handlers

import "net/http"

func CacheControl(cache string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", cache)
		h.ServeHTTP(w, r)
	})
}
