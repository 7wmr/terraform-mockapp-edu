#!/bin/bash

echo "Download and install Go binaries"
curl -O https://storage.googleapis.com/golang/go1.11.8.linux-amd64.tar.gz

tar -xf go1.11.8.linux-amd64.tar.gz && rm go1.11.8.linux-amd64.tar.gz
mv go /usr/local/

echo "Set enviornment variables required for Go"
export GOROOT=/usr/local/go
export GOPATH=$HOME/go

cat <<'EOF' >> ~/.bashrc
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
EOF

. ~/.bashrc

echo "Installing package - gin gonic"
go get "github.com/gin-gonic/gin"
echo "Installing package - go uuid"
go get "github.com/satori/go.uuid"


