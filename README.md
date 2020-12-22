# gordo

## What does it do?
This is a small program that will check your ticket numbers against the [EL PAÍS API](https://servicios.elpais.com/sorteos/loteria-navidad/api/) and output the prices won with them

## How to use it?
To use it just edit the name(s) and number(s) in `tickets.json`

## How to run it?
Just clone this repository and run it using `go run .` or `go build -ldflags="-s -w" -trimpath .` and `./gordo`

Alternatively you can download pre-built binaries from the [release page](https://github.com/thunderstorm99/gordo/releases)
