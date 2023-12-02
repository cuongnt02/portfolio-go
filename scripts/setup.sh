#!/bin/bash

# syncing database
cmd=$1

if [[($cmd == 'pull')]]; then
    heroku pg:pull postgresql-metric-69532 notetaker-ntc02 --app notetaker-ntc02
fi

if [[($cmd == "push")]]; then
    heroku pg:reset --confirm notetaker-ntc02
    heroku pg:push notetaker-ntc02 postgresql-metric-69532 --app notetaker-ntc02
fi
