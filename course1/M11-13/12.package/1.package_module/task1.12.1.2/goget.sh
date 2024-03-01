#!/bin/bash
go mod init $1
if [ $# -eq 0 ]; then
  echo "Необходимо указать имя модуля"
  exit
fi
go get github.com/yuin/goldmark