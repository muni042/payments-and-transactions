package client

import (
	"log"
	pb "sample/api/customers/v1"

	"google.golang.org/grpc"
)

// customer client
func NewCustomersClient() (pb.CustomersServiceClient, func(), error) {

	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := conn.Close(); err != nil {
			log.Println("failed to close grpc connection:", err)
		}
	}

	client := pb.NewCustomersServiceClient(conn)
	return client, cleanup, nil
}
