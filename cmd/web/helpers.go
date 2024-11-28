package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n", err.Error())
	app.errorLog.Output(2, trace)

	app.render(w, http.StatusInternalServerError, "error.html", &AsciiArt{
		ErrorText: http.StatusText(http.StatusInternalServerError),
		ErrorCode: http.StatusInternalServerError,
	})
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	app.render(w, status, "error.html", &AsciiArt{
		ErrorText: http.StatusText(status),
		ErrorCode: status,
	})
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *AsciiArt) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}
