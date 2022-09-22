package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/muhaidil13/booking/pkg/config"
	"github.com/muhaidil13/booking/pkg/handlers"
	"github.com/muhaidil13/booking/pkg/render"
)

const portNum = ":8080"

var App config.Appconfig
var session *scs.SessionManager

func main() {

	// Ganti ini jika diproducttion
	App.InProduct = false

	session = scs.New()
	session.Lifetime = 10 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = App.InProduct

	App.Session = session
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

	// http.HandleFunc("/", handlers.Repo.Home)
	fmt.Println("Listen on Port ", portNum)

	// http.ListenAndServe(portNum, nil)
	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(&App),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
