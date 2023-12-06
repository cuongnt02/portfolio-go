package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"notetaker.ntc02.net/internal/models"
	"notetaker.ntc02.net/internal/validator"
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
    data.FormActionPath = "/notes/create"

    app.render(w, http.StatusOK, "note-create.html", data)

}


type noteCreateForm struct {
    Title string `form:"title"`
    Content string `form:"content"`
    validator.Validator `form:"-"`
}



func (app *application) noteCreatePost(w http.ResponseWriter, r *http.Request) {

    var form noteCreateForm

    err := app.decodePostForm(r, &form)
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank.")
    form.CheckField(validator.MaxChars(form.Title, 200), "title", "This field cannot be more than 200 characters long.")
    form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank.")

    if !form.Valid() {
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

    app.sessionManager.Put(r.Context(), "flash", "Note Created Successfully!")

    http.Redirect(w, r, fmt.Sprintf("/notes/view/%d", id), http.StatusSeeOther)

}

func (app *application) noteEdit(w http.ResponseWriter, r *http.Request) {

    param := httprouter.ParamsFromContext(r.Context())

    id, err := strconv.Atoi(param.ByName("id"))
    if err != nil {
        app.notFound(w)
        return
    }
    
    note, err := app.notes.Get(id)
    if err != nil {
        if errors.Is(err, models.ErrNoRecord) {
            app.notFound(w)
            return
        } else {
            app.serverError(w, err)
            return
        }
    }

    data := app.newTemplateData(r)
    data.FormActionPath = fmt.Sprintf("/notes/update/%d", id)
    form := noteCreateForm {
        Title: note.Title,
        Content: note.Content,
    }

    data.Form = form
    app.render(w, http.StatusOK,  "note-create.html", data)

}

func (app *application) noteUpdatePost(w http.ResponseWriter, r *http.Request) {

    param := httprouter.ParamsFromContext(r.Context())

    id, err := strconv.Atoi(param.ByName("id"))
    if err != nil {
        app.notFound(w)
        return
    }
    
    var form noteCreateForm
    err = app.decodePostForm(r, &form)
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }


    form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank.") 
    form.CheckField(validator.MaxChars(form.Title, 200), "title", "This field cannot be more than 200 characters long.") 
    form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank.") 

    if !form.Valid() {
        data := app.newTemplateData(r)
        data.Form = form
        app.render(w, http.StatusUnprocessableEntity, "note-create.html", data)
        return
    }

    id, err = app.notes.Update(form.Title, form.Content, id)
    if err != nil {
        app.serverError(w, err)
        return
    }

    app.sessionManager.Put(r.Context(), "flash", "Note Updated Successfully!")
    
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

type userSignupForm struct {
    Username string `form:"username"`
    Password string `form:"password"`
    Email string `form:"email"`
    validator.Validator `form:"-"`
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Display Login Form Here...")
}

func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "End user's session Here")
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {

    data := app.newTemplateData(r)
    data.FormActionPath = "/user/signup"

    data.Form = userSignupForm{}
    
    app.render(w, http.StatusOK, "user-signup.html", data)

}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Process Login Form Here")
}

func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Process Signup Form Here")
}

