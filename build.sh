#!/bin/bash

APP_NAME=golang
echo "Building $APP_NAME ..."
APP_VERSION=v2
APP_TAGS=$APP_NAME:$APP_VERSION

docker build -t $APP_TAGS .

docker run -d \
-p 8088:8083 \
--name $APP_TAGS \
--network=go-net \
-v go-app:/usr/share/awesomeProject \
$APP_TAGS

echo "Building finished successfully"