#!/usr/bin/sh 

echo "LP:: HELLO postdeploy"

# Kill the service again to make sure it's the latest version running in the end.
# If there was a request between the old version stopped and the new started by EB then
# systemd would start another old version, which is killed here.
# We kill it anyway and the next request will start the server again.
[[ "active" == $(systemctl is-active web.service) ]] && systemctl stop web.service

systemctl status web.service
