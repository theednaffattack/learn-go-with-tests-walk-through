# Notes

## How to Run Tests

simple form: `go test`
recursive testing of sub-folders: `go test ./...`
recursive testing of sub-folders (verbose): `go test ./... -v`

## Installing GoDoc

Installation: `go get golang.org/x/tools/cmd/godoc`

Solution adapted from:
https://stackoverflow.com/a/63442485/9448010

`export PATH="$PATH:/usr/local/go/bin"`
`export PATH="$PATH:$HOME/go/bin"`

## Run Benchmarks

`go test -bench=.`
