package biz
import (
	"context"

	
	pb "cards/api/cards/v1"
	"cards/internal/helper"
	"cards/internal/mapper"
)
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