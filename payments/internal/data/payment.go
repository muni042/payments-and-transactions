package data

import (
	"context"
	"payments/internal/model"
	"payments/snowflake"

	"gorm.io/gorm"
)

type PaymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

// create
func (pr *PaymentRepo) InsertDB(ctx context.Context, p *model.Payment) error {
	p.PaymentID = snowflake.NextSnowflakeID()
	return pr.db.WithContext(ctx).Create(p).Error
}
func (pr *PaymentRepo) GetDB(
	ctx context.Context,
	id int64,
) (*model.Payment, error) {

	var p model.Payment

	err := pr.db.
		WithContext(ctx).
		Where("payment_id = ?", id).
		First(&p).
		Error

	if err != nil {
		return nil, err
	}

	return &p, nil
}
func (pr *PaymentRepo) UpdateDB(
	ctx context.Context,
	p *model.Payment,
) error {

	result := pr.db.
		WithContext(ctx).
		Model(&model.Payment{}).
		Where("payment_id = ?", p.PaymentID).
		Updates(p)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
