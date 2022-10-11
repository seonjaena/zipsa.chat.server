@ECHO off

cd %~dp0
cd ../

go mod init zipsa.chat.server

go get -u github.com/magiconair/properties
go get -u github.com/sirupsen/logrus
go get -u github.com/google/martian/log