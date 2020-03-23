# GKE-gRPC-Service-Ingress
In this repository you can find a simple Service implemeting a gRPC server and a REST gateway server. The aim of it is to deploy a service consuming both gRPC/REST behind an Ingress loadBalancer on GKE (Google Kubernetes Engine)

# Running the server
    go run main.go ENVIRONMENT=local NAMESPACE=default GRPC_PORT=8080 REST_PORT=8081

# Running the client (to test)
    cd .test_client && go run main.go SERVICE_ENDPOINT=localhost:8080 && cd ..