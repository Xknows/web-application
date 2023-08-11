package handlers

import (
	"net/http"
	"web/packeges/render"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "home.tmpl")
}
