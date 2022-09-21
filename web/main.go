package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/muhaidil13/web-golang/pkg/config"
	"github.com/muhaidil13/web-golang/pkg/handlers"
	"github.com/muhaidil13/web-golang/pkg/render"
)

const portNum = ":8080"

func main() {
	var App config.Appconfig

	// Mengambil Data Pointer dari package render bertipe *template.Template
	tc, err := render.RenderTemplate()
	if err != nil {
		log.Fatal(err)
	}
	// Kemudian Dimasukkan Kedalam App.TemplateCache
	// bertipe sama *template.Template
	App.TemplateCache = tc
	App.UseCache = false

	// handler Repo
	repo := handlers.NewRepo(&App)
	handlers.NewHandler(repo)

	// memasukkan data ke parameter pointer yang ada pada package render
	render.NewTemplate(&App)

	http.HandleFunc("/", handlers.Repo.Home)
	fmt.Println("Listen on Port ", portNum)
	http.ListenAndServe(portNum, nil)
}
