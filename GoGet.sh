#!/bin/sh
if [ $# -eq 1 ]; then
        GOPATH=$1
fi

cd signaling/cmd
go get -v 
