package data

import (
	"accounts/internal/model"
	"accounts/snowflake"
	"context"
	"time"

	"gorm.io/gorm"
)

type AccountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *AccountRepo {
	return &AccountRepo{db: db}
}
//create
func (ar *AccountRepo) InsertDB(ctx context.Context, a *model.Account) error {
	a.AccountID = snowflake.NextSnowflakeID()
	return ar.db.WithContext(ctx).Create(a).Error
}
//get
func (ar *AccountRepo) GetDB(ctx context.Context, id int64) (*model.Account, error) {
	var a model.Account
	if err := ar.db.WithContext(ctx).Where("account_id=?", id).First(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}
//update
func (ar *AccountRepo) UpdateDB(ctx context.Context, a *model.Account) error {
	return ar.db.WithContext(ctx).Model(&model.Account{}).Where("account_id=?", a.AccountID).Updates(a).Error
}
func (ar *AccountRepo) DeleteDB(ctx context.Context, id int64) error {
	now := time.Now()

	return ar.db.WithContext(ctx).
		Model(&model.Account{}).
		Where("account_id = ?", id).
		Updates(map[string]interface{}{
			"discarded":    "Y",
			"discard_time": now,
		}).Error
}
//existbyid
func (r *AccountRepo) ExistsByID(ctx context.Context, id int64) (bool, error) {

	var count int64

	err := r.db.WithContext(ctx).
		Model(&model.Account{}).
		Where("account_id = ?", id).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
//debit
func (ar *AccountRepo) Debit(
	ctx context.Context,
	accountID int64,
	amount float64,
) (bool, error) {

	result := ar.db.WithContext(ctx).
		Model(&model.Account{}).
		Where(
			"account_id = ? AND available_balance >= ?",
			accountID,
			amount,
		).
		Update(
			"available_balance",
			gorm.Expr("available_balance - ?", amount),
		)

	if result.Error != nil {
		return false, result.Error
	}

	// No rows affected means:
	// - account not found OR
	// - insufficient balance
	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
//credit
func (ar *AccountRepo) Credit(
	ctx context.Context,
	accountID int64,
	amount float64,
) error {

	return ar.db.WithContext(ctx).
		Model(&model.Account{}).
		Where("account_id = ?", accountID).
		Update(
			"available_balance",
			gorm.Expr("available_balance + ?", amount),
		).Error
}
