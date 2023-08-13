#!/bin/bash

set -x
# Remote server information
SERVER_USER="root"
SERVER_HOST="203.194.113.211"
SSH_PORT=22

# Define SSH private key
SSH_PRIVATE_KEY="~/.ssh/habi"

# Backup file name that created with clickhouse_backup_run_on_server.sh
BACKUP_FILE_NAME="/root/clickhouse_$(date '+%Y%m%d_%H%M%S').tgz"

# Project directory on the remote server
PROJECT_DIR="/root/dev/street"

# Where a backup file will be store in local machine
TARGET_LOCAL_DIR="../tmp"

# Run a remote script on a local machine over ssh
ssh -p $SSH_PORT -i $SSH_PRIVATE_KEY $SERVER_USER@$SERVER_HOST 'bash' $PROJECT_DIR/deploy/backup_clickhouse_on_server.sh

# Copy backup file from remote server to local machine
rsync --remove-source-files -r -e "ssh -p ${SSH_PORT} -i ${SSH_PRIVATE_KEY}" $SERVER_USER@$SERVER_HOST:$BACKUP_FILE_NAME $TARGET_LOCAL_DIR