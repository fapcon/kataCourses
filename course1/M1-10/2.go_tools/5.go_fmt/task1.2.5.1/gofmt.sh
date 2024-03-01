#!/bin/bash

# Проверяем, что передан аргумент
if [ -z "$1" ]; then
  echo "Usage: gofmt.sh <filename>"
  exit 1
fi

# Форматируем файл
go fmt $1