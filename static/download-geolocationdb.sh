#!/usr/bin/env bash

# source: https://ipapi.is/geolocation.html
aria2c https://ipapi.is/data/geolocationDatabaseIPv6.csv.zip
aria2c https://ipapi.is/data/geolocationDatabaseIPv4.csv.zip
