#!/bin/bash

cd $1
rm main
go build src/main.go
./main
