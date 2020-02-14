#!/bin/bash

VERSION=`cat version`

docker run -d -p 9345:9345 --restart always --name daycare-scheduler pokus2000/daycare-scheduler:$VERSION
