package render

import (
	"net/http"
	"testing"

	"github.com/muhaidil13/booking/internal/model"
)

func TestAddDefaultData(t *testing.T) {
	var td model.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("AddDefault Data Error")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplate = "./../../template"
	r, err := getSession()

	if err != nil {
		t.Error("Get Session Error")
	}
	var ww mywrite
	err = RenderFile(&ww, r, "home.page.html", &model.TemplateData{})
	if err != nil {
		t.Error("error renderfile")
	}
	// err = RenderFile(&ww, r, "homaaae.page.html", &model.TemplateData{})
	// if err != nil {
	// 	t.Error("error renderfile")
	// }
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/sxample", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}

func TestNewTemplate(t *testing.T) {
	NewTemplate(app)
}

func TestCreateTemplate(t *testing.T) {
	pathToTemplate = "./../../template"
	_, err := CreateTemplate()
	if err != nil {
		t.Error("error Create Template")
	}
}
