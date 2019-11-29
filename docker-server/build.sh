#!/bin/bash
echo "building"
docker build -t pigeoncai/grpc-server:1 .

echo "pushing"
docker push pigeoncai/grpc-server:1