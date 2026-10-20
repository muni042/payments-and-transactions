package mapper

import (
	"strconv"
	"time"

	pb "cards/api/cards/v1"
	"cards/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Create mapper
func CreateReqToCardBiz(req *pb.CreateCardRequest) *model.Card {

	expiry, _ := time.Parse("2006-01-02", req.ExpiryDate)
	daily, _ := strconv.ParseFloat(req.DailyLimit, 64)
	monthly, _ := strconv.ParseFloat(req.MonthlyLimit, 64)

	return &model.Card{
		AccountID:      req.AccountId,
		CardNumberHash: req.CardNumberHash,
		CardType:       req.CardType,
		CardStatus:     "active",
		ExpiryDate:     expiry,
		DailyLimit:     daily,
		MonthlyLimit:   monthly,
		PinAttempts:    0,
		Pin:            req.Pin,
	}
}

// Update mapper
func UpdateReqToCardBiz(req *pb.UpdateCardRequest) *model.Card {

	daily, _ := strconv.ParseFloat(req.DailyLimit, 64)
	monthly, _ := strconv.ParseFloat(req.MonthlyLimit, 64)

	return &model.Card{
		CardID:       req.CardId,
		CardStatus:   req.CardStatus,
		DailyLimit:   daily,
		MonthlyLimit: monthly,
		PinAttempts:  req.PinAttempts,
	}
}

// Model -> Proto
func BizToCardResponse(c *model.Card) *pb.CardResponse {

	res := &pb.CardResponse{
		CardId:         c.CardID,
		AccountId:      c.AccountID,
		CardNumberHash: c.CardNumberHash,
		CardType:       c.CardType,
		CardStatus:     c.CardStatus,
		ExpiryDate:     c.ExpiryDate.Format("2006-01-02"),
		DailyLimit:     strconv.FormatFloat(c.DailyLimit, 'f', -1, 64),
		MonthlyLimit:   strconv.FormatFloat(c.MonthlyLimit, 'f', -1, 64),
		PinAttempts:    c.PinAttempts,
	}

	if !c.CreatedAt.IsZero() {
		res.CreatedAt = timestamppb.New(c.CreatedAt)
	}
	if !c.LastUsedAt.IsZero() {
		res.LastUsedAt = timestamppb.New(c.LastUsedAt)
	}

	return res
}
