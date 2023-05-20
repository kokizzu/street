#!/usr/bin/env bash

mv /home/street/.env.production /home/street/.env &&
cp /home/street/street.service /lib/systemd/system/ &&
systemctl daemon-reload &&
systemctl enable street &&
systemctl restart street
