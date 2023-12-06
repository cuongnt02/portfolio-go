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
    router.Handler(http.MethodGet, "/.well-known/acme-challenge/*filepath", fileServer)

    dynamic := alice.New(app.sessionManager.LoadAndSave)


    router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
    router.Handler(http.MethodGet, "/notes", dynamic.ThenFunc(app.noteViewAll))
    router.Handler(http.MethodGet, "/notes/create", dynamic.ThenFunc(app.noteCreate))
    router.Handler(http.MethodPost, "/notes/create", dynamic.ThenFunc(app.noteCreatePost))
    router.Handler(http.MethodGet, "/notes/view/:id", dynamic.ThenFunc(app.noteView))
    router.Handler(http.MethodGet, "/notes/update/:id", dynamic.ThenFunc(app.noteEdit))
    router.Handler(http.MethodPost, "/notes/update/:id", dynamic.ThenFunc(app.noteUpdatePost))
    router.Handler(http.MethodGet, "/about", dynamic.ThenFunc(app.about))
    router.Handler(http.MethodGet, "/games", dynamic.ThenFunc(app.game))
    router.Handler(http.MethodGet, "/games/view", dynamic.ThenFunc(app.gameView))
    router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLogin))
    router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))
    router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignup))
    router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignupPost))
    router.Handler(http.MethodPost, "/user/logout", dynamic.ThenFunc(app.userLogoutPost))
    

    standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders, contentTypeHeaders)

    return standard.Then(router)
}
