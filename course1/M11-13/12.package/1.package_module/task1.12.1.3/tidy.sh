#!/bin/bash
go mod init $1
go get github.com/yuin/goldmark
go mod tidy