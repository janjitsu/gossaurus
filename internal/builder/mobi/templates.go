package mobi

import (
	"embed"
	"fmt"
	"io"
	"text/template"
)

// Embed all template files in the templates folder
//
//go:embed templates/*
var templateFS embed.FS

// Parse all templates
var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.ParseFS(templateFS, "templates/*")
	if err != nil {
		panic(fmt.Sprintf("Failed to parse templates: %v", err))
	}
}

// RenderTemplate executes a specific template with data
func RenderTemplate(fs io.Writer, templateName string,
	data interface{}) error {
	// Ensure template exists
	if tmpl.Lookup(templateName) == nil {
		return fmt.Errorf("template %s not found", templateName)
	}

	// Execute the template
	return tmpl.ExecuteTemplate(fs, templateName, data)
}
