package handlers

import (
	"net/http"

	"github.com/muhaidil13/web-golang/pkg/config"
	"github.com/muhaidil13/web-golang/pkg/render"
)

// send data Tamplate
type TemplateData struct {
	Stringmap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Datamap   map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

// repo use by handler
var Repo *Repository

// Repository Type
type Repository struct {
	App *config.Appconfig
}

// New Repo Create New Repository handler
func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

// set repository for hanlder
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderFile(w, "home.page.html", &TemplateData)
}
