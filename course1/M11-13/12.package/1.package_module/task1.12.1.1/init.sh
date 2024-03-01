#!/bin/bash
go mod init $1
if [ $# -eq 0 ]; then
  echo "Module name argument is missing"
  exit
fi