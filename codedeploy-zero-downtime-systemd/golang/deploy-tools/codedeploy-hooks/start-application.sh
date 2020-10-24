#!/bin/bash

set -ex

echo -e "Starting the application..."

systemctl enable web.socket
systemctl start web.socket

systemctl enable web.service
systemctl start web.service

echo -e "Started the application."
