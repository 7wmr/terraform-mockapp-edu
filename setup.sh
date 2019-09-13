#!/bin/bash

sudo add-apt-repository -y ppa:longsleep/golang-backports
sudo apt-get -y update
sudo apt-get -y install golang-go

go get "github.com/gin-gonic/gin"
go get "github.com/satori/go.uuid"


