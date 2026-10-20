package data

import (
	"context"

	"transactions/internal/model"
	"transactions/snowflake"

	"gorm.io/gorm"
)

type TransactionRepo struct {
	db *gorm.DB
}

//Constructor

func NewTransactionRepo(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

//CREATE

func (r *TransactionRepo) Create(
	ctx context.Context,
	tx *model.Transaction,
) error {

	tx.TransactionID = snowflake.NextSnowflakeID()

	return r.db.
		WithContext(ctx).
		Create(tx).
		Error
}

//READ – Get by Transaction ID

func (r *TransactionRepo) GetByID(
	ctx context.Context,
	transactionID int64,
) (*model.Transaction, error) {

	var tx model.Transaction

	err := r.db.
		WithContext(ctx).
		Where("transaction_id = ?", transactionID).
		First(&tx).
		Error

	if err != nil {
		return nil, err
	}

	return &tx, nil
}

//READ – List by Account ID

func (r *TransactionRepo) ListByAccountID(
	ctx context.Context,
	accountID int64,
) ([]*model.Transaction, error) {

	var txs []*model.Transaction

	err := r.db.
		WithContext(ctx).
		Where("account_id = ?", accountID).
		Order("created_on DESC").
		Find(&txs).
		Error

	if err != nil {
		return nil, err
	}

	return txs, nil
}

/*
EXISTS – Check if transaction exists by ID
*/
func (r *TransactionRepo) ExistsByID(
	ctx context.Context,
	transactionID int64,
) (bool, error) {

	var count int64

	err := r.db.
		WithContext(ctx).
		Model(&model.Transaction{}).
		Where("transaction_id = ?", transactionID).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

/*
UPDATE – Update transaction fields
*/
func (r *TransactionRepo) Update(
	ctx context.Context,
	tx *model.Transaction,
) error {

	result := r.db.
		WithContext(ctx).
		Model(&model.Transaction{}).
		Where("transaction_id = ?", tx.TransactionID).
		Updates(tx)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

/*
UPDATE – Update only status
*/
func (r *TransactionRepo) UpdateStatus(
	ctx context.Context,
	transactionID int64,
	status string,
) error {

	return r.db.
		WithContext(ctx).
		Model(&model.Transaction{}).
		Where("transaction_id = ?", transactionID).
		Update("status", status).
		Error
}

/*
DELETE – Hard delete
*/
func (r *TransactionRepo) Delete(
	ctx context.Context,
	transactionID int64,
) error {

	return r.db.
		WithContext(ctx).
		Where("transaction_id = ?", transactionID).
		Delete(&model.Transaction{}).
		Error
}

/*
DELETE – Soft archive (recommended)
*/
func (r *TransactionRepo) Archive(
	ctx context.Context,
	transactionID int64,
) error {

	return r.db.
		WithContext(ctx).
		Model(&model.Transaction{}).
		Where("transaction_id = ?", transactionID).
		Update("archived_at", gorm.Expr("NOW()")).
		Error
}
