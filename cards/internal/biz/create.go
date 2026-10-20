package biz
import (
	"context"

	accountspb "accounts/api/accounts/v1"
	pb "cards/api/cards/v1"
	"cards/internal/helper"
	"cards/internal/mapper"
	"fmt"
)
func (uc *CardUsecase) Create(
	ctx context.Context,
	req *pb.CreateCardRequest,
) (*pb.CardResponse, error) {

	//Validate card request
	if err := helper.ValidateCreateCard(req); err != nil {
		return nil, err
	}

	//CHECK ACCOUNT EXISTS (Accounts service via gRPC)
	existsResp, err := uc.acc.CheckAccountExists(
		ctx,
		&accountspb.GetAccountRequest{
			AccountId: req.AccountId,
		},
	)
	if err != nil {
		return nil, err
	}

	if !existsResp.Exists {
		return nil, fmt.Errorf("account %d does not exist", req.AccountId)
	}

	//Map proto model
	card := mapper.CreateReqToCardBiz(req)

	//Insert DB
	if err := uc.repo.Create(ctx, card); err != nil {
		return nil, err
	}

	//Map modelproto
	return mapper.BizToCardResponse(card), nil
}
