#!/bin/sh -e
ulimit -n $(ulimit -n -H)

export CADDYPATH=$SNAP_COMMON

environment=$(snapctl get environment)
if [ -n "$environment" ]; then
    export $environment
fi

email=$(snapctl get email)
if [ -n "$email" ]; then
    email_option="-email $email"
fi

snapctl get caddyfile | $SNAP/bin/caddy -agree $email_option -conf stdin