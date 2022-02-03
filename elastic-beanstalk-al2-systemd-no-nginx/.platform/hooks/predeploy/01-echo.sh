#!/usr/bin/sh 

echo "HELLO predeploy"

# Handle the web service and nginx before enabling systemd socket activation.
# [[ "active" == $(systemctl is-active web.service) ]] && systemctl stop web.service
# [[ "active" == $(systemctl is-active nginx.service) ]] && systemctl stop nginx.service

# Start the systemd socket!
cp ./etc/systemd/system/web.socket /etc/systemd/system/web.socket
systemctl enable web.socket
systemctl start web.socket

systemctl status web.socket
