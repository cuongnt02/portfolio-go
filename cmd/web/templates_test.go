package main

import (
	"testing"
	"time"

	"notetaker.ntc02.net/internal/assert"
)

func TestHumanDate(t *testing.T) {

    tests := []struct {
        name string
        tm time.Time
        want string
    } {
        {
            name: "UTC",
            tm : time.Date(2023, 12, 10, 03, 26, 0, 0, time.UTC),
            want: "Dec 10 2023 at 03:26", 
        },
        {
            name: "Empty",
            tm : time.Time{},
            want: "", 
        },
        {
            name: "CET",
            tm : time.Date(2023, 12, 10, 03, 26, 0, 0, time.FixedZone("CET", 1*60*60)),
            want: "Dec 10 2023 at 02:26", 
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            hd := humanDate(tt.tm)

            assert.Equal(t, hd, tt.want)
        })
    }


}
