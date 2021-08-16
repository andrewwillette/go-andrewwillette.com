#!/usr/bin/env bash

API_PORT=6969
REACT_PORT=80
starting_dir=$(pwd)

# build go executable in production version and start serving on port 6969
cd ../backend || return
go build -ldflags "-s -w"
go run . &
cd "$starting_dir" || return

# build react executable and start serving (using nginx) on port 80
cd ../frontend || return
npm run build
npm run start-prod