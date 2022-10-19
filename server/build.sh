#/bin/bash
export GOOS=darwin
export GOARCH=amd64

go build -o '../static-server-mac'


export GOOS=linux
export GOARCH=amd64
go build -o '../static-server-linux'
