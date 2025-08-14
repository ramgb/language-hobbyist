#!/bin/bash

cd $1
go build src/main.go
./main
