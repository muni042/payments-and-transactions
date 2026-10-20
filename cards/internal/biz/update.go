package biz
import (
	"context"

	pb "cards/api/cards/v1"
	"cards/internal/helper"
	"cards/internal/mapper"

)
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