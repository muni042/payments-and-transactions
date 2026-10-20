package biz

import (
	"context"

	pb "cards/api/cards/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (uc *CardUsecase) VerifyPin(
	ctx context.Context,
	req *pb.VerifyCardPinRequest,
) (*pb.VerifyCardPinResponse, error) {

	card, err := uc.repo.VerifyPinByAccountID(
		ctx,
		req.AccountId,
		req.Pin,
	)
	if err != nil {
		// real DBerror
		return nil, status.Error(
			codes.Internal,
			"database error",
		)
	}

	if card.CardStatus == "blocked" {
		return nil, status.Error(
			codes.PermissionDenied, // 403
			"card blocked",
		)
	}

	if card.PinAttempts > 0 {
		return nil, status.Error(
			codes.PermissionDenied,
			"invalid pin",
		)
	}

	//Correct PIN
	return &pb.VerifyCardPinResponse{
		Verified:    true,
		PinAttempts: 0,
		CardStatus:  card.CardStatus,
	}, nil
}
