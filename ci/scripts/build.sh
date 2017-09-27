#!/bin/bash

cd $WORKDIR
GOOS=linux GOARCH=amd64 go build -v -o ../concourse-demo-bin/concourse-demo-linux-amd64
GOOS=darwin GOARCH=amd64 go build -v -o ../concourse-demo-bin/concourse-demo-darwin-amd64
