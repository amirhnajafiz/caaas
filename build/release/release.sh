#!/bin/bash

echo "Terminating" $2;
echo "Release" $1;

sudo docker stop authx:$2;
sudo docker build . -f build/package/Dockerfile -t authx:$1;
sudo docker run -d -p 8080:8080 authx:$1;

sudo docker ps;
