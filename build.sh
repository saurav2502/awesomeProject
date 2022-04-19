#!/bin/bash

APP_NAME=golang
echo "Building $APP_NAME ..."
APP_VERSION=v5
APP_TAGS=$APP_NAME:$APP_VERSION

if [ $(arch) == "x86_64" ]; then
  echo "os architecture: $(arch)"
  docker build -t "saurav2502/$APP_TAGS" -f Dockerfile_arm .
else
  docker build -t "saurav2502/$APP_TAGS" .
fi

docker build -t "saurav2502/$APP_TAGS" .

docker run -d \
  -p 9005:9004 \
  --name $APP_NAME"-"$APP_VERSION \
  --network=go-net \
  -v go-app:/usr/share/awesomeProject \
  "saurav2502/$APP_TAGS"

echo "Building finished successfully"
