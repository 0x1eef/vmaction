#!/usr/bin/env bash

##
# vars
max=200
attempt=1

##
# main
sleep 10
virsh list --all
virsh domiflist testvm
#virsh net-list --all
#virsh net-dumpxml default

#while [ $attempt -le $max ]; do
#    IP=$(virsh domifaddr testvm | sed -n 's/.*ipv4 *\([0-9.]*\)\/.*/\1/p')
#    if [ -n "$IP" ]; then
#        echo -n "$IP"
#        exit 0
#    fi
#    sleep 1
#    attempt=$((attempt + 1))
#done
exit 1
