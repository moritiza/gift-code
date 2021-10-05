package repository

import (
	"github.com/moritiza/gift-code/discount-service/app/dto"
	"github.com/moritiza/gift-code/discount-service/app/entity"
	"gorm.io/gorm"
)

type UsedDiscountRepository interface {
	Create(report entity.UsedDiscount) *gorm.DB
	GetByCodeAndMobile(d dto.SubmitDiscount) (dto.SubmitDiscount, *gorm.DB)
}

// usedDiscountRepository satisfy UsedDiscountRepository interface
type usedDiscountRepository struct {
	db *gorm.DB
}

// NewUsedDiscountRepository creates a new used discount repository with the given dependencies
func NewUsedDiscountRepository(db *gorm.DB) UsedDiscountRepository {
	return &usedDiscountRepository{
		db: db,
	}
}

// Create do insert operation on used_discounts table and return database result
func (udr *usedDiscountRepository) Create(usedDiscount entity.UsedDiscount) *gorm.DB {
	r := udr.db.Model(entity.UsedDiscount{}).Create(&usedDiscount)
	return r
}

// GetByCodeAndMobile do read operation on used_discounts table and find used discount by discount code and mobile
func (udr *usedDiscountRepository) GetByCodeAndMobile(d dto.SubmitDiscount) (dto.SubmitDiscount, *gorm.DB) {
	var u dto.SubmitDiscount

	r := udr.db.Model(entity.UsedDiscount{}).Where("code = ?", d.Code).Where("mobile = ?", d.Mobile).First(&u)
	return u, r
}
