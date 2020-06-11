package handlers

import (
	"github.com/drag0ns3c/httpg/sys"
	"github.com/drag0ns3c/httpg/templates"
	"html/template"
	"net/http"
)

func SystemInfoHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.New("")
		tmpl.Parse(templates.Index)
		err := tmpl.Execute(w, sys.New())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
