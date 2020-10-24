#!/bin/bash

set -ex

echo -e "Stopping the application..."

UNIT_EXISTS=$( systemctl list-units | grep web )
[[ ! -z "$UNIT_EXISTS" ]] && systemctl stop web.service

echo -e "Stopped the application."
