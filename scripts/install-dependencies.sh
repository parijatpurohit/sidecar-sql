#!/bin/bash

which -s brew
if [[ $? != 0 ]] ; then
    # Install Homebrew
    echo "brew not installed. installing now"
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
else
    echo "brew installed.. checking for updates"
    brew update
fi

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
brew install protobuf
brew install protoc-gen-go

go get gopkg.in/yaml.v2 github.com/go-sql-driver/mysql github.com/jinzhu/gorm google.golang.org/grpc golang.org/x/net github.com/golang/protobuf

./generate.sh go