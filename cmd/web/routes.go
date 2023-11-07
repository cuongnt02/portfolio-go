package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)
func (app *application) routes() http.Handler {

    router := httprouter.New()

    router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        app.notFound(w)
    })


    fileServer := http.FileServer(http.Dir("./ui/static/"))
    router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))


    router.HandlerFunc(http.MethodGet, "/", app.home)
    router.HandlerFunc(http.MethodGet, "/notes", app.noteViewAll)
    router.HandlerFunc(http.MethodGet, "/notes/create", app.noteCreate)
    router.HandlerFunc(http.MethodPost, "/notes/create", app.noteCreatePost)
    router.HandlerFunc(http.MethodGet, "/notes/view/:id", app.noteView)
    router.HandlerFunc(http.MethodGet, "/about", app.about)

    defaultChain := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

    return defaultChain.Then(router)
}
