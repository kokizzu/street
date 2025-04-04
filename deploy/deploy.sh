#!/usr/bin/env bash

SERVER=root@benalu.dev
./no_restart.sh &&
ssh $SERVER sudo bash /home/street/reload_services.sh

echo '

 kill and restart street process manually on byobu sudo su - street
 since systemd is broken there

   '