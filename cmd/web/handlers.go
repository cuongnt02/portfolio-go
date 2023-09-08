package main

import (
	"fmt"
	"net/http"
	"strconv"
    "errors"

	"jforum.cuongnt02.org/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }
    
    data := app.newTemplateData(r)

    app.render(w, http.StatusOK, "index.html", data)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {

    data := app.newTemplateData(r)

    app.render(w, http.StatusOK, "about.html", data)

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

    http.Redirect(w, r, fmt.Sprintf("/threads/view?id=%d", id), http.StatusSeeOther)

}

func (app * application) viewThread(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }

    thread, err := app.threads.Get(id)


    if err != nil {

        if errors.Is(err, models.ErrNoRecord) {
            app.notFound(w)
        } else {
            app.serverError(w, err)
        }
        return
    }

    data := app.newTemplateData(r)
    data.Thread = thread

    app.render(w, http.StatusOK, "topic.html", data)


}

func (app *application) viewAllThread(w http.ResponseWriter, r *http.Request) {
    threads, err := app.threads.GetAll()

    if err != nil {
        app.serverError(w, err)
        return
    }

    data := app.newTemplateData(r)
    data.Threads = threads
    
    app.render(w, http.StatusOK, "threads.html", data)

}
