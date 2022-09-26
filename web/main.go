package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/muhaidil13/booking/internal/config"
	"github.com/muhaidil13/booking/internal/handlers"
	"github.com/muhaidil13/booking/internal/helpers"
	"github.com/muhaidil13/booking/internal/model"
	"github.com/muhaidil13/booking/internal/render"
)

const portNum = ":8080"

var App config.Appconfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := Run()
	if err != nil {
		log.Fatal(err)
	}

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

func Run() error {
	// Store data di session
	gob.Register(model.Reservation{})

	// Ganti ini jika diproducttion
	App.InProduct = false

	// Menulis error pada terminal
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	App.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.LstdFlags)
	App.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 10 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = App.InProduct

	App.Session = session
	// Mengambil Data Pointer dari package render bertipe *template.Template
	tc, err := render.CreateTemplate()
	if err != nil {
		log.Fatal(err)
		return err
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
	helpers.NewHelpers(&App)

	// Set Helpers data
	return nil
}
