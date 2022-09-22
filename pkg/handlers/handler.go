package handlers

import (
	"net/http"

	"github.com/muhaidil13/booking/pkg/config"
	"github.com/muhaidil13/booking/pkg/model"
	"github.com/muhaidil13/booking/pkg/render"
)

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
	StringMap := make(map[string]string)
	StringMap["users"] = "hello users"
	StringMap["test"] = "hello user"
	render.RenderFile(w, "home.page.html", &model.TemplateData{
		StringMap: StringMap,
	})
}
