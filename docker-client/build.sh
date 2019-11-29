#!/bin/bash
echo "building"
docker build -t pigeoncai/grpc-client:1 .

echo "pushing"
docker push pigeoncai/grpc-client:1