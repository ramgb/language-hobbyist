#!/bin/bash

# Create a new project directory
mkdir -p $1
cd $1

mkdir -p src

touch src/main.go

# Create a new go.mod file
go mod init $1

# Create a new main.go file
