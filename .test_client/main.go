package main

import (
	"context"
	pb "github.com/Emixam23/GKE-gRPC-Service-Ingress/interface"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"strings"
	"time"
)

// TestServiceClient is used to test
type TestServiceClient struct {

}

func main() {

	testServiceClient := TestServiceClient{}

	serviceEndpoint := testServiceClient.GetProgramArguments()

	// Set up a threads service stub connection to the service.
	serviceConn, err := grpc.Dial(serviceEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	service := pb.NewTestServiceClient(serviceConn)

	log.Println()
	log.Println("---- TEST ----")
	testServiceClient.HealthCheck(service)
	testServiceClient.HelloWorld(service, "Emixam23")
	log.Println("---- TEST ----")
	log.Println()

}

func (s *TestServiceClient) GetProgramArguments() (serviceEndpoint string) {

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

func (s *TestServiceClient) HelloWorld(conn pb.TestServiceClient, name string) {

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

func (s *TestServiceClient) HealthCheck(conn pb.TestServiceClient) {

	// Registering the instance to the dispatcher
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	if _, err := conn.HealthCheck(ctx, &pb.HealthCheckRequest{}); err != nil {
		s, _ := status.FromError(err)
		log.Fatalln(s.Message())
	} else {
		log.Println("200 OK")
	}

}