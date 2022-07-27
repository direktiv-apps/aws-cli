#!/bin/sh

docker build -t aws-cli . && docker run -p 9191:8080 aws-cli