package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/muhaidil13/booking/pkg/config"
	"github.com/muhaidil13/booking/pkg/model"
)

var app *config.Appconfig

var functions = template.FuncMap{}

// initialisiasi value app
func NewTemplate(e *config.Appconfig) {
	app = e
}

func AddDefaultData(e *model.TemplateData) *model.TemplateData {
	return e
}

func RenderFile(w http.ResponseWriter, form string, data *model.TemplateData) {

	// Development Mode
	var rt map[string]*template.Template
	if app.UseCache {
		rt = app.TemplateCache
	} else {
		rt, _ = RenderTemplate()
	}

	// akan menerima value t value, ok true jika ada dan false jika tidak ada
	t, ok := rt[form]
	if !ok {
		log.Fatal("Tidak terdapat data")
	}
	buf := new(bytes.Buffer)

	data = AddDefaultData(data)

	// Exekusi dan Melempar data
	_ = t.Execute(buf, data)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Terdapat error saat menulis", err)
	}

}

// remder file Html sebagai map
func RenderTemplate() (map[string]*template.Template, error) {
	mycache := map[string]*template.Template{}
	v, err := filepath.Glob("./template/*.page.html")
	if err != nil {
		return mycache, err
	}
	for _, page := range v {
		name := filepath.Base(page)
		fmt.Println("yang dijalankan adalah ", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return mycache, err
		}
		mathes, err := filepath.Glob("./template/*.layout.html")
		if err != nil {
			return mycache, err
		}
		if len(mathes) > 0 {
			ts, err := ts.ParseGlob("./template/*.layout.html")
			if err != nil {
				return mycache, err
			}
			mycache[name] = ts
		}
	}
	return mycache, nil

}
