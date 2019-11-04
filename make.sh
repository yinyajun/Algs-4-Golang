#!/usr/bin/env bash

if [ "$1" == "clean" ]; then
    rm bin/* pkg/* -rf
    exit
fi

cd $(dirname $0) &&

CURDIR=`pwd`
OLDGOPATH=$GOPATH
export GOPATH=$CURDIR
echo $GOPATH

echo "formatting code..."
$GOROOT/bin/go fmt algs4

echo "install algs4"
$GOROOT/bin/go install algs4
if [ $? -ne ""0 ]; then
    echo "install algs4 failed"
    exit -1
fi


export GOPATH=$OLDGOPATH
echo "finished"