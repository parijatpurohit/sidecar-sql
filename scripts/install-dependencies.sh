#!/bin/bash

which -s go
if [[ $? != 0 ]] ; then
  echo "go not installed. installing now"
  brew install go
  export GOPATH=~/go
  export PATH=$PATH:$GOPATH
fi

echo "current path is: "
pwd
echo $GOPATH

go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u gopkg.in/yaml.v2 github.com/go-sql-driver/mysql github.com/jinzhu/gorm google.golang.org/grpc golang.org/x/net github.com/golang/protobuf

export PATH=$PATH:/$GOPATH/bin

./generator.sh go