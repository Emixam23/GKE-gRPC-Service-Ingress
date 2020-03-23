#!/usr/bin/env bash

docker build -t gcr.io/<PROJECT_ID>/GKEgRPCService \
    -t gcr.io/<PROJECT_ID>/GKEgRPCService:latest \
    . \
    --build-arg ENVIRONMENT=ingresstest \
    --build-arg NAMESPACE=default \
    --build-arg GRPC_PORT=8080 \
    --build-arg REST_PORT=8081 \
&& docker push gcr.io/<PROJECT_ID>/GKEgRPCService:latest \
&& kubectl apply -f .kubernetes/service.yaml --force