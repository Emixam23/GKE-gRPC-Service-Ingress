# Dockerfile References: https://docs.docker.com/engine/reference/builder/

############################################ MULTI STAGE BUILD PART 1 ##############################################

# Start from golang v1.11 base image
FROM golang:1.14 as builder

# Add Maintainer Info
LABEL maintainer="Maxime GUITTET <maxime.guittet@say-eyes.com>"

# Setting env variable
ENV GO111MODULE=on

# Creating/Cd work directory
WORKDIR /service

# Copying module files
COPY go.mod .
COPY go.sum .

# RUN --mount=type=ssh go mod download
RUN go mod download

# Copying sources
COPY . .

# Run go build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

RUN chmod +x main

############################################ MULTI STAGE BUILD PART 2 ##############################################

# Start from alpine image
FROM alpine

# Creating work directory
WORKDIR /service

# Copy the certificats and executable into new Docker image
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /service/main /service/
RUN apk add --no-cache curl

## Get required ARGs and put them into ENVs variables
ARG ENVIRONMENT
ARG NAMESPACE
ARG GRPC_PORT
ARG REST_PORT
ENV _ENVIRONMENT=ingresstest${ENVIRONMENT}
ENV _NAMESPACE=default${NAMESPACE}
ENV _GRPC_PORT=${GRPC_PORT}
ENV _REST_PORT=${REST_PORT}

# Expose port
EXPOSE ${GRPC_PORT}
EXPOSE ${REST_PORT}

ENTRYPOINT ./main "ENVIRONMENT=${_ENVIRONMENT}" "NAMESPACE=${_NAMESPACE}" "GRPC_PORT=${_GRPC_PORT}" "REST_PORT=${_REST_PORT}"
