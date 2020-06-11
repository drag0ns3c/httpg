package handlers

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"
	"net/http"
)

func StaticHandler(box *packr.Box) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		asset, err := box.FindString(vars["asset"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(asset))
	})
}
