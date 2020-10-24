#!/bin/bash

set -ex

echo -e "Validate service stage"

PUBLIC_IP=$( curl http://169.254.169.254/latest/meta-data/public-ipv4 )

for i in `seq 5`; do
    curl -v http://127.0.0.1:80/lambros && echo -e "Server validated local IP."
    curl -v "http://${PUBLIC_IP}/lambros" && echo -e "Server validated public IP." && exit 0

    echo -e "Validator sleeping for 1 second..."
    sleep 1
done

echo -e "Validator failed!"

exit 1
