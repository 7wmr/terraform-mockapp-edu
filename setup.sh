#!/bin/bash

wget https://storage.googleapis.com/golang/go1.11.8.linux-amd64.tar.gz

tar -xf go1.11.8.linux-amd64.tar.gz && rm go1.11.8.linux-amd64.tar.gz
mv go /usr/local && mkdir -p /usr/local/gopath

echo "set enviornment variables required for Go"
export GOROOT=/usr/local/go
export GOPATH=/usr/local/gopath
cat <<EOF >> ~/.bashrc
export GOROOT=/usr/local/go
export GOPATH=/usr/local/gopath
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
EOF

. ~/.bashrc

echo "Installing package - gin gonic"
go get "github.com/gin-gonic/gin"
echo "Installing package - go uuid"
go get "github.com/satori/go.uuid"


