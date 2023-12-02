package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/lib/pq"

	"notetaker.ntc02.net/internal/models"
)


type application struct {
    errorLog    *log.Logger
    infoLog     *log.Logger
    notes       *models.NoteModel
    templateCache map[string] *template.Template
    formDecoder *form.Decoder
    sessionManager *scs.SessionManager
}

func OpenDB(dsn string) (*sql.DB, error) {
    
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    return db, nil
}


func main() {
    
    addr := ":" +  os.Getenv("PORT")
    if addr == ":" {
        addr = ":8000"
    }

    dsn := os.Getenv("DATABASE_URL")


    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    db, err := OpenDB(dsn)
    if err != nil {
        errorLog.Fatal(err);
    }

    defer db.Close()

    templateCache, err := newTemplateCache()
    if err != nil {
        errorLog.Fatal(err)
    }

    formDecoder := form.NewDecoder()

    sessionManager := scs.New()
    sessionManager.Store = postgresstore.New(db)
    sessionManager.Lifetime = 12 * time.Hour
    sessionManager.Cookie.Secure = true

    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
        notes: &models.NoteModel{DB: db},
        templateCache: templateCache,
        formDecoder: formDecoder,
        sessionManager: sessionManager,
    }



    srv := &http.Server {
        Addr: addr,
        ErrorLog: errorLog,
        Handler: app.routes(),
    }


    infoLog.Printf("Starting server on: %s", addr)
    err = srv.ListenAndServe()
    errorLog.Fatal(err)

}


