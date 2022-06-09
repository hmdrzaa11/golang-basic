package config

import (
	"html/template"
	"log"
)

// AppConfig contains global configuration for entire app
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InfoLog       *log.Logger
}
