#!/bin/sh

docker build -t aws-cli . && docker run -e DIREKTIV_TEST=true -p 9191:8080 -p 9292:9292 aws-cli