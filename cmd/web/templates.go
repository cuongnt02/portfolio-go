package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"notetaker.ntc02.net/internal/models"
	"notetaker.ntc02.net/ui"
)

type templateData struct {
    Note *models.Note
    Notes []*models.Note
    CurrentYear int
    FormActionPath string
    Form any
    Flash string
    IsAuthenticated bool
    CSRFToken string
}

func newTemplateCache() (map[string] *template.Template, error) {
    cache := map[string] *template.Template{}

    
    pages, err := fs.Glob(ui.Files,"html/pages/*.html")
    if err != nil {
        return nil, err
    }

    for _, page := range pages {
        
        name := filepath.Base(page)

        patterns := []string {
            "html/base.html",
            "html/partials/*.html",
            page,
        }

        ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
        if err != nil {
            return nil, err
        }


        cache[name] = ts

    }

    return cache, nil

}

func humanDate(t time.Time) string {
    if t.IsZero() {
        return ""
    }

    return t.UTC().Format("Jan 02 2006 at 15:04")
}


var functions = template.FuncMap {
    "humanDate" : humanDate,
}

