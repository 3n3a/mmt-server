#!/bin/bash

(
    cd src/server
    GOARCH=arm64 GOOS=linux go build -o server .
)

mv src/server/server system/bin/server