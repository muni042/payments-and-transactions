package clientcard

import (
	pb1 "cards/api/cards/v1"
	"log"

	"google.golang.org/grpc"
)

func NewCardsClient() (pb1.CardsServiceClient, func(), error) {

	conn, err := grpc.Dial("localhost:9200", grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := conn.Close(); err != nil {
			log.Println("failed to close grpc connection:", err)
		}
	}

	client := pb1.NewCardsServiceClient(conn)
	return client, cleanup, nil
}
