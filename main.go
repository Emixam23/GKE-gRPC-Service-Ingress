package main

import (
	"cloud.google.com/go/profiler"
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/Emixam23/GKE-gRPC-Service-Ingress/interface"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

const SERVICE = "test-service"

// main is the starting point of the current micro service.
func main() {

	// Initialising the service server
	serviceServer := &TestServiceServer{

	}

	// Program arguments
	environment, namespace, gRPCPort, RESTPort := serviceServer.GetProgramArguments()

	// Profiler initialization, best done as early as possible.
	// Only Available on GCP.
	if environment != "local" {
		// local-test-service
		// dev-Max-test-service
		// preprod-test-service
		service := environment
		if namespace != "" {
			service = service + "-" + namespace
		}
		service = service + "-" + SERVICE

		// Profiler initialization, best done as early as possible.
		if err := profiler.Start(profiler.Config{
			Service:   strings.ToLower(service),
			ProjectID: "<PROJECT_ID>",
		}); err != nil {
			log.Fatalf("Unable to start the profiler: [%s].\n", err.Error())
		}
	}

	// Just some logs
	log.Println("Try running service ...")

	// grpc server
	go serviceServer.RunAsGRPCServer(fmt.Sprintf("0.0.0.0:%d", gRPCPort))

	// grpc gateway server
	// Run as gRPC gateway server
	if err := serviceServer.RunAsGRPCGatewayServer(fmt.Sprintf("0.0.0.0:%d", gRPCPort), fmt.Sprintf("0.0.0.0:%d", RESTPort)); err != nil {
		glog.Fatal(err)
	}

}

//
// Info & Arguments
//

func (s *TestServiceServer) GetProgramArguments() (environment string, namespace string, gRPCPort uint, RESTPort uint) {

	args := os.Args[1:]
	if len(os.Args) > 1 {
		for _, arg := range args {
			s := strings.Split(arg, "=")
			switch s[0] {
			case "ENVIRONMENT":
				environment = s[1]
				break
			case "NAMESPACE":
				namespace = s[1]
				break
			case "GRPC_PORT":
				if port, err := strconv.Atoi(s[1]); err != nil {
					log.Fatalf("Invalid %s provided: [%s]\n", s[0], s[1])
				} else {
					gRPCPort = uint(port)
				}
				break
			case "REST_PORT":
				if port, err := strconv.Atoi(s[1]); err != nil {
					log.Fatalf("Invalid %s provided: [%s]\n", s[0], s[1])
				} else {
					RESTPort = uint(port)
				}
				break
			default:
				log.Fatalf("unknown argument provided: %s", arg)
			}
		}
	} else {
		if ENVIRONMENT := os.Getenv("ENVIRONMENT"); ENVIRONMENT != "" {
			environment = ENVIRONMENT
		}
		if NAMESPACE := os.Getenv("NAMESPACE"); NAMESPACE != "" {
			namespace = NAMESPACE
		}
		if GRPC_PORT := os.Getenv("GRPC_PORT"); GRPC_PORT != "" {
			if port, err := strconv.Atoi(GRPC_PORT); err != nil {
				log.Fatalf("Invalid %s provided: [%s]\n", "GRPC_PORT", GRPC_PORT)
			} else {
				gRPCPort = uint(port)
			}
		}
		if REST_PORT := os.Getenv("REST_PORT"); REST_PORT != "" {
			if port, err := strconv.Atoi(REST_PORT); err != nil {
				log.Fatalf("Invalid %s provided: [%s]\n", "REST_PORT", REST_PORT)
			} else {
				RESTPort = uint(port)
			}
		}
	}

	////////////////////////////////////////////////////////////////////////
	// Here each parameter which requires to be checked should be checked //
	////////////////////////////////////////////////////////////////////////

	// Checking ports (grpc, rest)
	if gRPCPort == 0 {
		log.Fatalf("Invalid gRPC port provided [%d]. The gRPC port must be > 0.\n", gRPCPort)
	} else if RESTPort == 0 {
		log.Fatalf("Invalid REST port provided [%d]. The REST port must be > 0.\n", RESTPort)
	} else if gRPCPort == RESTPort {
		log.Fatalf("Invalid ports provided [%d/%d]. They must be different.\n", gRPCPort, RESTPort)
	}


	log.Printf("ENVIRONMENT=ingresstest%s\n", environment)
	log.Printf("NAMESPACE=default%s\n", namespace)
	log.Printf("GRPC_PORT=%d\n", gRPCPort)
	log.Printf("REST_PORT=%d\n", RESTPort)

	return environment, namespace, gRPCPort, RESTPort
}

//
// gRPC
//

func (s *TestServiceServer) RunAsGRPCServer(gRPCAddr string) {

	lis, err := net.Listen("tcp", gRPCAddr) // gRPC
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("gRPC Listening on %s\n", gRPCAddr)

	server := grpc.NewServer(s.withServerUnaryInterceptor())
	pb.RegisterTestServiceServer(server, s)
	// Register reflection service on gRPC server.
	reflection.Register(server)

	if server.Serve(lis) != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *TestServiceServer) withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(s.serverInterceptor)
}

// Authorization unary interceptor function to handle authorize per RPC call
func (s *TestServiceServer) serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now().UTC()

	// Getting the meta data
	_, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	// Calls the handler
	response, err := handler(ctx, req)
	if err != nil {
		// Logging if there is an error
		s, _ := status.FromError(err)
		log.Printf("Request - Method:%s\tDuration:%s\tError:%s\n",
			info.FullMethod,
			time.Since(start),
			s.Message())
	} else {
		// Logging for the response
		r, _ := json.Marshal(response)
		log.Printf("Request - Method:%s\tDuration:%s\tResponse:%s\n",
			info.FullMethod,
			time.Since(start),
			string(r))
	}

	return response, err
}

//
// gRPC Gateway
//

func (s *TestServiceServer) RunAsGRPCGatewayServer(gRPCAddr string, RESTAddr string) error {

	opts := []grpc.DialOption{grpc.WithInsecure()}
	mux := runtime.NewServeMux()

	if err := pb.RegisterTestServiceHandlerFromEndpoint(context.Background(), mux, gRPCAddr, opts); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
	log.Printf("HTTP Listening on %s\n", RESTAddr)

	return http.ListenAndServe(RESTAddr, mux)

}

//
// Implementation
//

// TestServiceServer is used to implement TestServiceServer.TestServiceServer.
type TestServiceServer struct {

}

func (s *TestServiceServer) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{}, nil
}

func (s *TestServiceServer) HelloWorld(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{Content:fmt.Sprintf("HelloWorld to you %s!", in.Name)}, nil
}