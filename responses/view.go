package responses

import (
	"html/template"
	"net/http"
)

type view struct {
	template *template.Template
	data     interface{}
}

type ViewData struct {
	Title string
	Data  interface{}
}

func (v view) Write(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	v.template.Execute(w, v.data)
}

func View(t *template.Template, data ViewData) Response {
	return view{t, data}
}
