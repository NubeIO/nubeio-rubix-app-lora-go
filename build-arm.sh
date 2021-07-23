#!/bin/bash
env GOOS=linux GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc GOARM=7 go build -ldflags="-extldflags=-static" -tags sqlite_omit_load_extension
