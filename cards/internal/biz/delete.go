package biz
import (
	"context"

	
	pb "cards/api/cards/v1"
	"cards/internal/helper"
	"cards/internal/mapper"
)
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
