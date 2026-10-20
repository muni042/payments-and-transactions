package data

import (
	"context"
	"sample/internal/model"
	"sample/snowflake"
	"time"

	"errors"

	"gorm.io/gorm"
)

type CustomerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{db: db}
}

// create
func (cr *CustomerRepo) InsertDB(ctx context.Context, c *model.Customer) error {
	now := time.Now()
	c.CustomerID = snowflake.NextSnowflakeID()
	c.CreationTime = &now
	c.ModifiedTime = &now
	c.Discarded = ptr("N")
	return cr.db.WithContext(ctx).Create(c).Error
}

// get
func (cr *CustomerRepo) GetDB(ctx context.Context, id int64) (*model.Customer, error) {
	var c model.Customer
	err := cr.db.WithContext(ctx).Where("customer_id=?", id).First(&c).Error
	return &c, err
}

// update
func (cr *CustomerRepo) UpdateDB(ctx context.Context, c *model.Customer) error {
	return cr.db.WithContext(ctx).
		Model(&model.Customer{}).
		Where("customer_id=?", c.CustomerID).
		Updates(c).Error
}

// delete
func (cr *CustomerRepo) DeleteDB(ctx context.Context, id int64) error {
	now := time.Now()

	result := cr.db.WithContext(ctx).
		Model(&model.Customer{}).
		Where("customer_id = ? AND status != ?", id, "closed").
		Updates(map[string]interface{}{
			"status":       "closed",
			"discarded":    "Y",
			"discard_time": now,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("customer account already closed or not found")
	}
	return nil
}

func ptr(s string) *string {
	return &s
}

// existbyid
func (cr *CustomerRepo) ExistsByID(ctx context.Context, id int64) (bool, error) {

	var count int64

	err := cr.db.WithContext(ctx).
		Model(&model.Customer{}).
		Where("customer_id = ? AND discarded = 'N'", id).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
