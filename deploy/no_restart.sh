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
  --exclude='etc_clickhouse-server' \
  --exclude='opt_tarantool' \
  --exclude='var_lib_tarantool' \
  --exclude='tmp_docker-mailserver' \
  --exclude='var_lib_clickhouse' \
  --exclude='var_log_letsencrypt' \
  --exclude='var_log_mail' \
  --exclude='var_mail' \
  --exclude='var_mail-state' \
  --exclude='tmp_docker-mailserver' \
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
upx $BIN_NAME &&
rsync -avz --progress . $SERVER:/home/street

# cp ../.env.override . &&