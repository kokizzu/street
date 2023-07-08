#!/usr/bin/env bash

SERVER=root@benalu.dev
VERSION=$(date +%Y.%m.%d)-$(git rev-parse --short HEAD)$(git diff --quiet || echo '-dev')

BIN_NAME=street.exe

( cd .. &&
GOOS=linux GOARCH=amd64 go build -o $BIN_NAME \
    -ldflags="-X main.VERSION='$VERSION'" ) &&
rsync -apvz ../static ./ &&
mv ../$BIN_NAME . &&
rsync --delete -a \
  --exclude='_*' \
  --exclude='.*' \
  --exclude='tmp' \
  --exclude='*.svelte' \
  --exclude='*.sh' \
  --exclude='*.md' \
  --exclude='package*.json' \
  --exclude='tsconfig.json' \
  --exclude='svelte.config.js' \
  --exclude='build.js' \
  --exclude='node_modules' \
  --include='*.ico' \
  --include='*.png' \
  --include='*.jpg' \
  --include='*.svg' \
  --include='*.js' \
  --include='*.html' \
  --include='*.css' \
  ../svelte/ svelte &&
cp ../.env.override . &&
upx $BIN_NAME &&
rsync -avz --progress . $SERVER:/home/street &&
ssh $SERVER sudo bash /home/street/reload_services.sh

