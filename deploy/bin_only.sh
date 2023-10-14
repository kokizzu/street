#!/usr/bin/env bash

SERVER=root@benalu.dev
VERSION=$(date +%Y.%m.%d)-$(git rev-parse --short HEAD)$(git diff --quiet || echo '-dev')

BIN_NAME=street_bin.exe

( cd .. &&
GOOS=linux GOARCH=amd64 go build -o $BIN_NAME \
    -ldflags="-X main.VERSION='$VERSION'" ) &&
mv ../$BIN_NAME . &&
upx $BIN_NAME &&
rsync -avz --progress $BIN_NAME $SERVER:/home/street/

