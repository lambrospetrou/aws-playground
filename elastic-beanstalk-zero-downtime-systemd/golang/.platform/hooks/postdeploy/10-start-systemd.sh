#!/bin/bash

echo -e "Copying and Running the systemd commands"

cp -r deploy-tools/systemd /etc/

systemctl enable web_2.socket
systemctl start web_2.socket

systemctl enable web_2.service
systemctl start web_2.service
