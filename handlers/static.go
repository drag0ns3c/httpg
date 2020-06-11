package handlers

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
	"path/filepath"
)

func StaticHandler(box *packr.Box) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		asset, err := box.FindString(vars["asset"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		ext := filepath.Ext(vars["asset"])
		contentType := mime.TypeByExtension(ext)
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(asset))
	})
}
