#!/bin/sh -e

ulimit -n $(ulimit -n -H)

export XDG_DATA_HOME=$SNAP_COMMON
export XDG_CONFIG_HOME=$SNAP_COMMON

$SNAP/bin/caddy run --resume
