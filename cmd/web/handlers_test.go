package main

import (
	"net/http"
	"testing"

	"notetaker.ntc02.net/internal/assert"
)

func TestPing(t *testing.T) {

    app := newTestApplication(t)

    ts := newTestServer(t, app.routes())
    defer ts.Close()

    code, _, body := ts.get(t, "/ping")

    assert.Equal(t, code, http.StatusOK)
    assert.Equal(t, body, "OK")

}

func TestNoteView(t *testing.T) {
    app := newTestApplication(t)

    ts := newTestServer(t, app.routes())
    defer ts.Close()

    tests := []struct {
        name string
        urlPath string
        wantCode int
        wantBody string
    } {
        {
            name: "Valid ID",
            urlPath: "/notes/view/1",
            wantCode: http.StatusOK,
            wantBody: "NOTHING HERE YET",
        },
        {
            name: "Non-existent ID",
            urlPath: "/notes/view/2",
            wantCode: http.StatusNotFound,
        },
        {
            name: "Negative ID",
            urlPath: "/notes/view/-2",
            wantCode: http.StatusNotFound,
        },
        {
            name: "Decimal ID",
            urlPath: "/notes/view/2.2",
            wantCode: http.StatusNotFound,
        },
        {
            name: "String ID",
            urlPath: "/notes/view/hello",
            wantCode: http.StatusNotFound,
        },
        {
            name: "Empty ID",
            urlPath: "/notes/view/",
            wantCode: http.StatusNotFound,
        },
    }

    for _, tt:= range tests {
        t.Run(tt.name, func(t *testing.T) {
            code, _, body := ts.get(t, tt.urlPath)

            assert.Equal(t, code, tt.wantCode)
            
            if tt.wantBody != "" {
                assert.StringContains(t, body, tt.wantBody)
            }
        })
    }

}
