package handlers

import (
	"github.com/drag0ns3c/httpg/sys"
	"github.com/gobuffalo/packr/v2"
	"html/template"
	"net/http"
)

func SystemInfoHandler(box *packr.Box) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		index, err := box.FindString("index.html")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tmpl := template.New("")
		tmpl.Parse(index)
		err = tmpl.Execute(w, sys.New())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
