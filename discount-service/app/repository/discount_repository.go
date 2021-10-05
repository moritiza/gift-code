package repository

import (
	"fmt"

	"github.com/moritiza/gift-code/discount-service/app/entity"
	"gorm.io/gorm"
)

type DiscountRepository interface {
	Create(report entity.Discount) *gorm.DB
	GetByCode(code string) (entity.Discount, *gorm.DB)
	IncrementCodeUsed(id uint64) *gorm.DB
}

// discountRepository satisfy DiscountRepository interface
type discountRepository struct {
	db *gorm.DB
}

// NewDiscountRepository creates a new discount repository with the given dependencies
func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	return &discountRepository{
		db: db,
	}
}

// Create do insert operation on discounts table and return database result
func (dr *discountRepository) Create(discount entity.Discount) *gorm.DB {
	r := dr.db.Model(entity.Discount{}).Create(&discount)
	return r
}

// GetByCode do read operation on discounts table and find discount by discount code
func (dr *discountRepository) GetByCode(code string) (entity.Discount, *gorm.DB) {
	var discount entity.Discount

	r := dr.db.Model(entity.Discount{}).Where("code = ?", code).First(&discount)
	return discount, r
}

// IncrementCodeUsed increment discount code used one unit
func (dr *discountRepository) IncrementCodeUsed(id uint64) *gorm.DB {
	r := dr.db.Exec(
		"UPDATE \"discounts\" SET used=used+1 WHERE id = " +
			fmt.Sprint(id) + " AND \"discounts\".\"deleted_at\" IS NULL",
	)
	return r
}
