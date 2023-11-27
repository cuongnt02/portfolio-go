package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/julienschmidt/httprouter"
	"notetaker.ntc02.net/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {


    data := app.newTemplateData(r)

    app.render(w, http.StatusOK, "home.html", data)

}

func (app *application) noteViewAll(w http.ResponseWriter, r *http.Request) {

    notes, err := app.notes.GetAll()
    if err != nil {
        app.serverError(w, err)
        return
    }

    data := app.newTemplateData(r)
    data.Notes = notes

    app.render(w, http.StatusOK, "notes.html", data)
}

func (app *application) noteView(w http.ResponseWriter, r *http.Request) {

    params := httprouter.ParamsFromContext(r.Context())

    id, err := strconv.Atoi(params.ByName("id"))
    if err != nil {
        app.notFound(w)
        return
    }

    
    note, err := app.notes.Get(id)
    if err != nil {
        if (errors.Is(err, models.ErrNoRecord)) {
            app.notFound(w)
        } else {
            app.serverError(w, err)
        }
        return
    }
    data := app.newTemplateData(r)
    data.Note = note

    app.render(w, http.StatusOK, "note-view.html", data)   
    
}

func (app *application) noteCreate(w http.ResponseWriter, r *http.Request) {

    data := app.newTemplateData(r)

    data.Form = noteCreateForm{}

    app.render(w, http.StatusOK, "note-create.html", data)

}


type noteCreateForm struct {
    Title string
    Content string
    FieldErrors map[string]string
}



func (app *application) noteCreatePost(w http.ResponseWriter, r *http.Request) {

    err := r.ParseForm()
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    form := noteCreateForm {
        Title: r.PostForm.Get("title"),
        Content: r.PostForm.Get("content"),
        FieldErrors: map[string]string{},
    }


    if strings.TrimSpace(form.Title) == "" {
        form.FieldErrors["title"] = "This field cannot be blank."
    } else if utf8.RuneCountInString(form.Title) > 200 {
        form.FieldErrors["title"] = "This field cannot be more than 200 characters long."
    }

    if strings.TrimSpace(form.Content) == "" {
        form.FieldErrors["content"] = "This field cannot be blank."
    }

    if len(form.FieldErrors) > 0 {
        data := app.newTemplateData(r)
        data.Form = form
        app.render(w, http.StatusUnprocessableEntity, "note-create.html", data)
        return
    }

    id, err := app.notes.Insert(form.Title, form.Content)
    if err != nil {
        app.serverError(w, err)
        return
    }

    http.Redirect(w, r, fmt.Sprintf("/notes/view/%d", id), http.StatusSeeOther)

}

func (app *application) about(w http.ResponseWriter, r *http.Request) {

    data := app.newTemplateData(r)

    app.render(w, http.StatusOK, "about.html", data)
}

func (app *application) game(w http.ResponseWriter, r *http.Request) {
    data := app.newTemplateData(r);

    app.render(w, http.StatusOK, "game.html", data)
}

func (app *application) gameView(w http.ResponseWriter, r *http.Request) {
    data := app.newTemplateData(r);
    app.render(w, http.StatusOK, "game-view.html", data)
}
