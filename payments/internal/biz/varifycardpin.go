package biz

import (
	cardspb "cards/api/cards/v1"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (uc *PaymentUsecase) verifyCardPin(
	ctx context.Context,
	accountID int64,
	pin string,
) error {

	resp, err := uc.cards.VerifyCardPin(ctx, &cardspb.VerifyCardPinRequest{
		AccountId: accountID,
		Pin:       pin,
	})
	if err != nil {
		return status.Error(codes.Unavailable, "card service unavailable")
	}

	log.Printf(
		"[VerifyCardPin] account_id=%d input_pin=%s verified=%v card_status=%s pin_attempts=%d",
		accountID,
		pin,
		resp.Verified,
		resp.CardStatus,
		resp.PinAttempts,
	)

	if resp.CardStatus == "blocked" {
		return status.Error(codes.PermissionDenied, "card blocked")
	}

	if !resp.Verified {
		return status.Error(codes.PermissionDenied, "invalid pin")
	}

	return nil
}
