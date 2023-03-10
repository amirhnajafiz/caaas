#!/bin/bash

echo "Terminating" $2;
echo "Release" $1;

sudo docker stop authx-container;
sudo docker build . -f build/package/Dockerfile -t authx:$1;
sudo docker run -d -p 8080:8080 authx:$1 --name authx-container;

sudo docker ps;
