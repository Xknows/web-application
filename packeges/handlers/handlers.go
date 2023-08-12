package handlers

import (
	"net/http"
	"web/packeges/render"
)

// call render and render specific template
func Home(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "home.page.tmpl")
}
