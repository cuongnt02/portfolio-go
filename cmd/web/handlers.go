package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }

    files := []string {
        "./ui/html/base.html",
        "./ui/html/partials/nav.html",
        "./ui/html/partials/footer.html",
        "./ui/html/pages/index.html",
    }

    ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        app.serverError(w, err)
    }
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {

    files := []string {
        "./ui/html/base.html",
        "./ui/html/partials/nav.html",
        "./ui/html/partials/footer.html",
        "./ui/html/pages/about.html",
    }

    ts, err := template.ParseFiles(files...)

    if err != nil {
        app.serverError(w, err)
        return
    }

    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        app.serverError(w, err)
    }

}

func (app *application) thread(w http.ResponseWriter, r*http.Request) {
    


    files := []string {
        "./ui/html/base.html",
        "./ui/html/partials/nav.html",
        "./ui/html/partials/footer.html",
        "./ui/html/pages/threads.html",
    }

    ts, err := template.ParseFiles(files...)

    if err != nil {
        app.serverError(w, err)
    }

    err = ts.ExecuteTemplate(w, "base", nil)

    if err != nil {
        app.serverError(w, err)
    }
}

func (app *application) createThread(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

    name := "Alcohol"

    id, err := app.threads.Insert(name)

    if err != nil {
        app.serverError(w, err)
        return
    }

    http.Redirect(w, r, fmt.Sprintf("/thread/view?id=%d", id), http.StatusSeeOther)

}
