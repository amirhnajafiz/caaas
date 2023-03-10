#!/bin/bash

echo "Terminating" $2;
echo "Release" $1;

docker stop authx:$2;
docker build . -f build/package/Dockerfile -t authx:$1;
docker run -d -p 8080:8080 authx:$1;
