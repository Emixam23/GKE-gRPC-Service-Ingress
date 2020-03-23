package main

import (
	"context"
	pb "github.com/Emixam23/GKE-gRPC-Service-Ingress/interface"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"strings"
	"time"
)

// GKEgRPCServiceClient is used to test
type GKEgRPCServiceClient struct {

}

func main() {

	testServiceClient := GKEgRPCServiceClient{}

	serviceEndpoint := testServiceClient.GetProgramArguments()

	// Set up a threads service stub connection to the service.
	serviceConn, err := grpc.Dial(serviceEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	service := pb.NewGKEgRPCServiceClient(serviceConn)

	log.Println()
	log.Println("---- TEST ----")
	testServiceClient.HealthCheck(serviceConn)
	testServiceClient.HelloWorld(service, "Emixam23")
	log.Println("---- TEST ----")
	log.Println()

}

func (s *GKEgRPCServiceClient) GetProgramArguments() (serviceEndpoint string) {

	args := os.Args[1:]
	if len(os.Args) > 1 {
		for _, arg := range args {
			s := strings.Split(arg, "=")
			switch s[0] {
			case "SERVICE_ENDPOINT":
				serviceEndpoint = s[1]
				break
			default:
				log.Fatalf("unknown argument provided: %s", arg)
			}
		}
	} else {
		if SERVICE_ENDPOINT := os.Getenv("SERVICE_ENDPOINT"); SERVICE_ENDPOINT != "" {
			serviceEndpoint = SERVICE_ENDPOINT
		}
	}

	////////////////////////////////////////////////////////////////////////
	// Here each parameter which requires to be checked should be checked //
	////////////////////////////////////////////////////////////////////////

	if serviceEndpoint == "" {
		log.Fatalf("Invalid service endpoint provided [%s]. It can't be empty.\n", serviceEndpoint)
	}

	log.Printf("SERVICE_ENDPOINT=%s\n", serviceEndpoint)

	return serviceEndpoint
}

func (s *GKEgRPCServiceClient) HelloWorld(conn pb.GKEgRPCServiceClient, name string) {

	// Registering the instance to the dispatcher
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	if response, err := conn.HelloWorld(ctx, &pb.HelloWorldRequest{
		Name: name,
	}); err != nil {
		s, _ := status.FromError(err)
		log.Fatalln(s.Message())
	} else {
		log.Println(response.Content)
	}

}

func (s *GKEgRPCServiceClient) HealthCheck(conn *grpc.ClientConn) {

	// Registering the instance to the dispatcher
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := healthpb.NewHealthClient(conn).Check(ctx, &healthpb.HealthCheckRequest{Service: "echo.EchoServer"})
	if err != nil {
		log.Fatalf("HealthCheck failed %+v\n", err)
	}

	if resp.GetStatus() != healthpb.HealthCheckResponse_SERVING {
		log.Fatalf("service not in serving state: %s\n", resp.GetStatus().String())
	}
	log.Printf("RPC HealthChekStatus:%v\n", resp.GetStatus())

}