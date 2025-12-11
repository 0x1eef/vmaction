#!/usr/bin/env bash

##
# vars
max=200
attempt=1

##
# main
while [ $attempt -le $max ]; do
    IP=$(sudo virsh domifaddr testvm | sed -n 's/.*ipv4 *\([0-9.]*\)\/.*/\1/p')
    if [ -n "$IP" ]; then
        echo -n "$IP"
        exit 0
    fi
    sleep 1
    attempt=$((attempt + 1))
done
exit 1
