#!/bin/bash
echo "Debug started..."
dlv debug $1
echo "Debug ended."
go build -o myprogram $1
dlv exec ./myprogram

$SHELL