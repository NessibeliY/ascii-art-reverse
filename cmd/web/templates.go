package main

import (
	"html/template"
	"path/filepath"
)

type AsciiArt struct {
	OrigText  string
	FinalText string
	ErrorText string
	ErrorCode int
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{
			"./templates/base.html",
			"./templates/partials/nav.html",
			page,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
