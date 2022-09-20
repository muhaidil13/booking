package handlers

import (
	"net/http"

	"github.com/muhaidil13/web-golang/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderFile(w, "home.page.html")
}
