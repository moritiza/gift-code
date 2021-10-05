package repository

import (
	"fmt"

	"github.com/moritiza/gift-code/wallet-service/app/dto"
	"github.com/moritiza/gift-code/wallet-service/app/entity"
	"gorm.io/gorm"
)

type CreditRepository interface {
	GetByMobile(mobile string) (entity.Credit, *gorm.DB)
	UpdateCreditByMobile(credit dto.SetDiscountCredit) *gorm.DB
	Create(credit entity.Credit) *gorm.DB
}

// creditRepository satisfy CreditRepository interface
type creditRepository struct {
	db *gorm.DB
}

// NewCreditRepository creates a new credit repository with the given dependencies
func NewCreditRepository(db *gorm.DB) CreditRepository {
	return &creditRepository{
		db: db,
	}
}

// GetByMobile do read operation on credits table and return founded credit with database result
func (cr *creditRepository) GetByMobile(mobile string) (entity.Credit, *gorm.DB) {
	var credit entity.Credit

	r := cr.db.Model(entity.Credit{}).Where("mobile = ?", mobile).First(&credit)
	return credit, r
}

// UpdateCreditByMobile do update operation on credits table and set credit by mobile and return database result
func (cr *creditRepository) UpdateCreditByMobile(credit dto.SetDiscountCredit) *gorm.DB {
	r := cr.db.Exec(
		"UPDATE \"credits\" SET credit=credit+" + fmt.Sprint(credit.Credit) + " WHERE mobile = '" +
			fmt.Sprint(credit.Mobile) + "' AND \"credits\".\"deleted_at\" IS NULL",
	)
	return r
}

// Create do insert operation on credits table and return database result
func (cr *creditRepository) Create(credit entity.Credit) *gorm.DB {
	r := cr.db.Model(entity.Credit{}).Create(&credit)
	return r
}
