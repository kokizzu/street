#!/usr/bin/env bash

server=root@benalu.dev
# TODO: execute box.backup.start()
rsync -apvz --progress $server:/home/street/var_lib_tarantool ../tmpdb/tarantool_backup
# TODO: execute box.backup.stop()