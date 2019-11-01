#!/usr/bin/env sh

while true; do
    /bin/flog "$@" >> $LOG_FILE_PATH
    sleep $SLEEP_TIME
done