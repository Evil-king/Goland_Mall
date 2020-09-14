#! /bin/bash
cd proto && protoc --go_out=plugins=grpc:../rpcservices InnerGameInfo.proto
protoc --go_out=plugins=grpc:../rpcservices Models.proto
cd ..