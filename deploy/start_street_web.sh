#!/usr/bin/env bash

mkdir -p `pwd`/logs
ofile=`pwd`/logs/access_`date +%Y%m%d%H%M%S`.log
echo "Starting server, logging to $ofile"
unbuffer time ./street.exe web 2>&1 | tee $ofile