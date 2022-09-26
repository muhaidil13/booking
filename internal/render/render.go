package render

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
	"github.com/muhaidil13/booking/internal/config"
	"github.com/muhaidil13/booking/internal/model"
)

var app *config.Appconfig

var pathToTemplate = "../template"

var functions = template.FuncMap{}

// initialisiasi value app
func NewTemplate(e *config.Appconfig) {
	app = e
}

// Set data For All Template
func AddDefaultData(e *model.TemplateData, r *http.Request) *model.TemplateData {

	e.Flash = app.Session.PopString(r.Context(), "flash")
	e.Error = app.Session.PopString(r.Context(), "error")
	e.Warning = app.Session.PopString(r.Context(), "warning")
	e.CSRFToken = nosurf.Token(r)
	return e
}

func RenderFile(w http.ResponseWriter, r *http.Request, form string, data *model.TemplateData) error {

	// Development Mode
	var rt map[string]*template.Template
	if app.UseCache {
		rt = app.TemplateCache
	} else {
		rt, _ = CreateTemplate()
	}

	// akan menerima value t value, ok true jika ada dan false jika tidak ada
	t, ok := rt[form]
	if !ok {
		log.Fatal("Tidak terdapat data")
		return errors.New("Cant get data form Tamplate")
	}
	buf := new(bytes.Buffer)

	data = AddDefaultData(data, r)

	// Exekusi dan Melempar data
	_ = t.Execute(buf, data)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Terdapat error saat menulis", err)
		return err
	}
	return nil

}

// remder file Html sebagai map
func CreateTemplate() (map[string]*template.Template, error) {
	mycache := map[string]*template.Template{}
	v, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplate))
	if err != nil {
		return mycache, err
	}
	for _, page := range v {
		name := filepath.Base(page)
		// fmt.Println("yang dijalankan adalah ", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return mycache, err
		}
		mathes, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", pathToTemplate))
		if err != nil {
			return mycache, err
		}
		if len(mathes) > 0 {
			ts, err := ts.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplate))
			if err != nil {
				return mycache, err
			}
			mycache[name] = ts
		}
	}
	return mycache, nil

}
