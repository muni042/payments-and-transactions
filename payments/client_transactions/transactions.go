package clienttransactions

import (
	"log"
	pb1 "transactions/api/transactions/v1"

	"google.golang.org/grpc"
)

func NewTransactionsClient() (pb1.TransactionServiceClient, func(), error) {

	conn, err := grpc.Dial("localhost:9300", grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := conn.Close(); err != nil {
			log.Println("failed to close grpc connection:", err)
		}
	}

	client := pb1.NewTransactionServiceClient(conn)
	return client, cleanup, nil
}
