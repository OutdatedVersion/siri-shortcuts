#!/bin/bash

# ./run.sh &
# strace -p<processID> -s9999 -e write

while true; do
    EVENT_HUB_URI="wss://home.outdatedversion.com/api/hub" ../build/client
    echo "client crashed with exit code $?. restart.." >&2
    sleep 1
done
