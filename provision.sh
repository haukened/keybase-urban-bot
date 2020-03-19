#!/usr/bin/env bash
keybase --no-auto-fork \
    --debug \
    oneshot \
    -u $KEYBASE_USERNAME \
    --paperkey "$(cat /run/secrets/$KEYBASE_USERNAME-paperkey)"
./app