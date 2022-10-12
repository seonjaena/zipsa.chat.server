#!/bin/bash

SCRIPT=$(readlink -f "$0")
SCRIPT_PATH=$(dirname "$SCRIPT")

cd $SCRIPT_PATH
cd ../

go mod init zipsa.chat.server

go get -u github.com/magiconair/properties
go get -u github.com/sirupsen/logrus
go get -u github.com/google/martian/log
go get -u github.com/streadway/amqp