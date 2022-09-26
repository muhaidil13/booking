package handlers

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
	"github.com/muhaidil13/booking/internal/config"
	"github.com/muhaidil13/booking/internal/model"
	"github.com/muhaidil13/booking/internal/render"
)

var pathToTemplate = "./../../template"

var functions = template.FuncMap{}
var App config.Appconfig
var session *scs.SessionManager

func getRoute() http.Handler {
	// Store data di session
	gob.Register(model.Reservation{})

	// Ganti ini jika diproducttion
	App.InProduct = true

	session = scs.New()
	session.Lifetime = 10 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = App.InProduct

	App.Session = session
	// Mengambil Data Pointer dari package render bertipe *template.Template
	tc, err := RenderTestTemplate()
	if err != nil {
		log.Fatal(err)
	}

	// Kemudian Dimasukkan Kedalam App.TemplateCache
	// bertipe sama *template.Template
	App.TemplateCache = tc
	App.UseCache = true
	// handler Repo
	repo := NewRepo(&App)
	NewHandler(repo)

	// memasukkan data ke parameter pointer yang ada pada package render
	render.NewTemplate(&App)

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	// mux.Use(NoSurf)

	mux.Use(SessionLoad)
	mux.Get("/", repo.Home)
	mux.Get("/about", repo.About)
	mux.Get("/contact", repo.Contact)
	mux.Get("/room", repo.Room)

	mux.Get("/reservation", repo.Reservation)
	mux.Post("/reservation", repo.PostReservation)
	mux.Post("/reservation-json", repo.AvailabilityJson)

	mux.Get("/reservation-result", repo.ReservationSummary)
	mux.Post("/make-reservation", repo.PostMake_Reservation)
	mux.Get("/make-reservation", repo.Make_Reservation)
	mux.Get("/pemesannan", repo.Pemesannan)

	fileserver := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))
	return mux
}

func NoSurf(next http.Handler) http.Handler {
	crfHandler := nosurf.New(next)
	crfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   App.InProduct,
		SameSite: http.SameSiteLaxMode,
	})
	return crfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// remder file Html sebagai map
func RenderTestTemplate() (map[string]*template.Template, error) {
	mycache := map[string]*template.Template{}
	v, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplate))
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
