package config

import (
	"log"
	"text/template"
)

type Appconfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
