package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type Appconfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduct     bool
	Session       *scs.SessionManager
}
