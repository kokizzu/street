#!/usr/bin/env bash

SERVER=root@benalu.dev
./no_restart.sh &&
ssh $SERVER sudo bash /home/street/reload_services.sh

