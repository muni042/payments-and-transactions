package data

import (
	"cards/model"
	"context"

	"cards/snowflake"

	"fmt"

	"gorm.io/gorm"
)

type CardRepo struct {
	data *gorm.DB
}

func NewCardRepo(db *gorm.DB) *CardRepo {
	return &CardRepo{data: db}
}

// create
func (r *CardRepo) Create(ctx context.Context, m *model.Card) error {
	m.CardID = snowflake.Generate()
	fmt.Println("sucess data")
	return r.data.WithContext(ctx).Create(m).Error

}

// getbyid
func (r *CardRepo) GetByID(ctx context.Context, id int64) (*model.Card, error) {
	var m model.Card
	err := r.data.WithContext(ctx).
		Where("card_id = ?", id).
		First(&m).Error
	return &m, err
}

// update
func (r *CardRepo) Update(ctx context.Context, m *model.Card) error {
	return r.data.WithContext(ctx).
		Model(&model.Card{}).
		Where("card_id = ?", m.CardID).
		Updates(m).Error
}

// delete
func (r *CardRepo) Delete(ctx context.Context, id int64) error {
	return r.data.WithContext(ctx).
		Model(&model.Card{}).
		Where("card_id = ? AND discarded = 'N'", id).
		Updates(map[string]interface{}{
			"discarded":    "Y",
			"discarded_at": gorm.Expr("NOW()"),
		}).Error
}

const maxPinAttempts = 3

func (r *CardRepo) VerifyPinByAccountID(
	ctx context.Context,
	accountID int64,
	pin string,
) (*model.Card, error) {

	var card model.Card

	// Fetch card
	if err := r.data.WithContext(ctx).
		Where("account_id = ? AND discarded = 'N'", accountID).
		First(&card).Error; err != nil {
		return nil, err
	}

	// Stop immediately if card already blocked
	if card.CardStatus == "blocked" {
		return &card, nil
	}

	// Wrong PIN
	if card.Pin != pin {

		// Do not increment beyond max
		if card.PinAttempts >= maxPinAttempts {
			return &card, nil
		}

		card.PinAttempts++

		updates := map[string]interface{}{
			"pin_attempts": card.PinAttempts,
		}

		// Block exactly at max
		if card.PinAttempts == maxPinAttempts {
			updates["card_status"] = "blocked"
			card.CardStatus = "blocked"
		}

		if err := r.data.WithContext(ctx).
			Model(&model.Card{}).
			Where("card_id = ?", card.CardID).
			Updates(updates).Error; err != nil {
			return nil, err
		}

		return &card, nil
	}

	// Correct PIN → reset attempts
	if err := r.data.WithContext(ctx).
		Model(&model.Card{}).
		Where("card_id = ?", card.CardID).
		Update("pin_attempts", 0).Error; err != nil {
		return nil, err
	}

	card.PinAttempts = 0
	return &card, nil
}

// get by account id (used by biz to fetch latest card state)
func (r *CardRepo) GetByAccountID(
	ctx context.Context,
	accountID int64,
) (*model.Card, error) {

	var card model.Card

	err := r.data.WithContext(ctx).
		Where("account_id = ? AND discarded = 'N'", accountID).
		First(&card).Error

	if err != nil {
		return nil, err
	}

	return &card, nil
}
