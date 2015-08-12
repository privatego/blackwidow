#!/bin/bash

CUR_PATH=`pwd`
export GOPATH=$CUR_PATH
echo "Current GOPATH=$CUR_PATH"

go build main

