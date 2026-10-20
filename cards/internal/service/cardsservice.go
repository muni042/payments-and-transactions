package service

import (
	"context"
	"fmt"

	pb "cards/api/cards/v1"
	"cards/internal/biz"
)

type CardsServiceService struct {
	pb.UnimplementedCardsServiceServer
	uc *biz.CardUsecase
}

func NewCardsServiceService(uc *biz.CardUsecase) *CardsServiceService {
	return &CardsServiceService{uc: uc}
}

// CREATE
func (s *CardsServiceService) CreateCard(
	ctx context.Context,
	req *pb.CreateCardRequest,
) (*pb.CardResponse, error) {

	fmt.Println("success card service")
	return s.uc.Create(ctx, req)
}

// GET by card_id
func (s *CardsServiceService) GetCard(
	ctx context.Context,
	req *pb.GetCardRequest,
) (*pb.CardResponse, error) {

	return s.uc.Get(ctx, req)
}

// UPDATE
func (s *CardsServiceService) UpdateCard(
	ctx context.Context,
	req *pb.UpdateCardRequest,
) (*pb.CardResponse, error) {

	return s.uc.Update(ctx, req)
}

// DELETE
func (s *CardsServiceService) DeleteCard(
	ctx context.Context,
	req *pb.GetCardRequest,
) (*pb.DeleteCardResponse, error) {

	card, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteCardResponse{
		Message: "Card deleted successfully",
		Card:    card,
	}, nil
}

// ✅ NEW: VERIFY CARD PIN USING ACCOUNT ID
func (s *CardsServiceService) VerifyCardPin(
	ctx context.Context,
	req *pb.VerifyCardPinRequest,
) (*pb.VerifyCardPinResponse, error) {

	return s.uc.VerifyPin(ctx, req)
}
