#!/bin/bash

export RABBITMQ_ADDRESS=amqp://test:test@10.0.2.15.5672
touch /pids
LISTEN_ADDRESS=10.29.1.1:12345 STORAGE_ROOT=/tmp/1 nohup go run dataservice/dataservice.go &> /dev/null & echo $! >> ./pids \
LISTEN_ADDRESS=10.29.1.2:12345 STORAGE_ROOT=/tmp/2 nohup go run dataservice/dataservice.go &>/dev/null echo $! >> ./pids \
#& LISTEN_ADDRESS=10.29.1.3:12345 STORAGE_ROOT=/tmp/3 go run dataservice/dataservice.go \
#& LISTEN_ADDRESS=10.29.1.4:12345 STORAGE_ROOT=/tmp/4 go run dataservice/dataservice.go \
#& LISTEN_ADDRESS=10.29.1.5:12345 STORAGE_ROOT=/tmp/5 go run dataservice/dataservice.go \
#& LISTEN_ADDRESS=10.29.1.6:12345 STORAGE_ROOT=/tmp/6 go run dataservice/dataservice.go \
LISTEN_ADDRESS=10.29.2.1:12345  nohup go run apiservice/apiservice.go &>/dev/null echo $! >> ./pids \
LISTEN_ADDRESS=10.29.2.2:12345  nohup go run apiservice/apiservice.go &>/dev/null echo $! >> ./pids
