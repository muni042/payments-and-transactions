package client

import (
	pb "accounts/api/accounts/v1"
	"log"

	"google.golang.org/grpc"
)

func AccountsClient() (pb.AccountsServiceClient, func(), error) {

	conn, err := grpc.Dial("localhost:9100", grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := conn.Close(); err != nil {
			log.Println("failed to close grpc connection:", err)
		}
	}

	client := pb.NewAccountsServiceClient(conn)
	return client, cleanup, nil
}
