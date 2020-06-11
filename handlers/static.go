package handlers

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"
	"log"
	"mime"
	"net/http"
	"path/filepath"
)

func StaticHandler(box *packr.Box) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		assetName := vars["asset"]
		asset, err := box.FindString(vars["asset"])

		if err != nil {
			log.Printf("request for asset %s not found", assetName)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		log.Printf("serving asset %s", assetName)

		ext := filepath.Ext(assetName)
		contentType := mime.TypeByExtension(ext)
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Cache-Control", "no-store")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(asset))
	})
}
