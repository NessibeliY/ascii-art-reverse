package main

import (
	"asciiartweb/nyeltay/algaliyev/internal"
	"net/http"
	"os"
	"strings"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	app.render(w, http.StatusOK, "home.html", &AsciiArt{})
}

func (app *application) ascii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	text := r.Form.Get("body")
	choiceOfFont := r.Form.Get("asciiType")
	if choiceOfFont != "shadow.txt" && choiceOfFont != "standard.txt" && choiceOfFont != "thinkertoy.txt" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	cwd, _ := os.Getwd()
	path := internal.TrimCwd(cwd)
	cwd = path + "/internal/banner/" + choiceOfFont

	_, err := os.Stat(cwd)

	if os.IsNotExist(err) {
		app.render(w, http.StatusInternalServerError, "error.html", &AsciiArt{
			ErrorText: http.StatusText(http.StatusInternalServerError),
			ErrorCode: http.StatusInternalServerError,
		})
		return
	}

	text = strings.ReplaceAll(text, "\r", "")
	result, err := internal.Convert(text, choiceOfFont)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	data := &AsciiArt{
		OrigText:  text,
		FinalText: result,
	}

	app.render(w, http.StatusOK, "ascii-art.html", data)
}
