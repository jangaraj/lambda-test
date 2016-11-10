#!/bin/bash

go get -d github.com/eawsy/aws-lambda-go/...
docker run --rm -v $GOPATH:/go -v $PWD:/tmp eawsy/aws-lambda-go --nopackage
