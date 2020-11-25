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

mkdir -p $GOPATH/src/github.com/parijatpurohit
cd $GOPATH/src/github.com/parijatpurohit