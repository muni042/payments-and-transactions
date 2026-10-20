package biz

import (
	accountspb "accounts/api/accounts/v1"
	"cards/internal/data"
)

type CardUsecase struct {
	repo *data.CardRepo
	acc  accountspb.AccountsServiceClient
}

func NewCardUsecase(
	repo *data.CardRepo,
	acc accountspb.AccountsServiceClient,
) *CardUsecase {
	return &CardUsecase{
		repo: repo,
		acc:  acc,
	}
}

/*
// CREATE CARD
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

	//Map proto → model
	card := mapper.CreateReqToCardBiz(req)

	//Insert DB
	if err := uc.repo.Create(ctx, card); err != nil {
		return nil, err
	}

	//Map model → proto
	return mapper.BizToCardResponse(card), nil
}

// GET CARD
func (uc *CardUsecase) Get(ctx context.Context, req *pb.GetCardRequest) (*pb.CardResponse, error) {

	if err := helper.ValidateCardID(req.CardId); err != nil {
		return nil, err
	}

	card, err := uc.repo.GetByID(ctx, req.CardId)
	if err != nil {
		return nil, err
	}

	return mapper.BizToCardResponse(card), nil
}

// UPDATE CARD
func (uc *CardUsecase) Update(ctx context.Context, req *pb.UpdateCardRequest) (*pb.CardResponse, error) {

	if err := helper.ValidateUpdateCard(req); err != nil {
		return nil, err
	}

	card := mapper.UpdateReqToCardBiz(req)

	if err := uc.repo.Update(ctx, card); err != nil {
		return nil, err
	}

	card, _ = uc.repo.GetByID(ctx, req.CardId)

	return mapper.BizToCardResponse(card), nil
}

// DELETE CARD
func (uc *CardUsecase) Delete(ctx context.Context, req *pb.GetCardRequest) (*pb.CardResponse, error) {

	if err := helper.ValidateCardID(req.CardId); err != nil {
		return nil, err
	}

	// fetch card before soft delete
	card, err := uc.repo.GetByID(ctx, req.CardId)
	if err != nil {
		return nil, err
	}

	// soft delete
	if err := uc.repo.Delete(ctx, req.CardId); err != nil {
		return nil, err
	}

	return mapper.BizToCardResponse(card), nil
}

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
		// real DB/system error
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
			codes.PermissionDenied, // 403
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
*/
