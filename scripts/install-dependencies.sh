#!/bin/bash

pwd
echo $GOPATH


go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u gopkg.in/yaml.v2 github.com/go-sql-driver/mysql github.com/jinzhu/gorm google.golang.org/grpc golang.org/x/net github.com/golang/protobuf

export PATH=$PATH:/$GOPATH/bin

./generator.sh go