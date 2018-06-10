#!/bin/bash
export GOPATH=$(pwd)

go get github.com/gin-gonic/gin

go get github.com/spf13/cobra
go get golang.org/x/sync/errgroup
go get github.com/gobuffalo/packr

echo install github.com/gobuffalo/packr
go install github.com/gobuffalo/packr/packr
