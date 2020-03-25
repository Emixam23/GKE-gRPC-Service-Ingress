#!/bin/sh

GOPATH=$(env | grep GOPATH | grep -oe '[^=]*$');
GOBIN=$(env | grep GOBIN | grep -oe '[^=]*$');
GOSRC=$(env | grep GOSRC | grep -oe '[^=]*$');
export PATH=$PATH:${GOBIN}

protofile=interface.proto
api_descriptor=interface.pb

########################################################################################################################
################################################# proto generation #####################################################
########################################################################################################################

# Go
protoc ${protofile} --proto_path . --proto_path "${GOSRC}"/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --descriptor_set_out="${api_descriptor}" --plugin=protoc-gen-grpc-gateway="${GOBIN}"/protoc-gen-grpc-gateway --swagger_out=logtostderr=true:./ --go_out=plugins=grpc:./ --grpc-gateway_out=logtostderr=true:./

## Kotlin
PROTOC_GEN_KOTLIN_PATH="/usr/local/lib/grpc-kotlin/protoc-gen-grpc-kotlin.jar"
protoc ${protofile} --proto_path . --plugin=protoc-gen-grpc_kotlin=${PROTOC_GEN_KOTLIN_PATH} --proto_path "${GOSRC}"/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --java_out=. --grpc_kotlin_out=.
mv GKEgRPCService/* .
mv io/rouz/grpc/* .
rm -rf GKEgRPCService io

## Swift
protoc ${protofile} --proto_path . --proto_path "${GOSRC}"/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swift_out=. --swiftgrpc_out=Client=true,Server=false:.

## C#
PROTOC_GEN_CSHARP_PATH="/usr/local/lib/grpc-csharp/grpc_csharp_plugin"
protoc ${protofile} --proto_path . --plugin="protoc-gen-grpc="${PROTOC_GEN_CSHARP_PATH} --proto_path "${GOSRC}"/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --csharp_out=. --grpc_out=.