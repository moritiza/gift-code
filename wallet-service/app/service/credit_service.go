package service

import (
	"errors"

	"github.com/moritiza/gift-code/wallet-service/app/dto"
	"github.com/moritiza/gift-code/wallet-service/app/entity"
	"github.com/moritiza/gift-code/wallet-service/app/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CreditService interface {
	GetCredit(credit dto.GetCredit) (dto.GetCredit, error)
	SetDiscountCredit(credit dto.SetDiscountCredit) (dto.SetDiscountCredit, error)
}

// creditService satisfy CreditService interface
type creditService struct {
	logger           logrus.Logger
	creditRepository repository.CreditRepository
}

// NewCreditService creates a new credit service with the given dependencies
func NewCreditService(l logrus.Logger, cr repository.CreditRepository) CreditService {
	return &creditService{
		logger:           l,
		creditRepository: cr,
	}
}

// GetCredit find user credit by mobile
func (cs *creditService) GetCredit(credit dto.GetCredit) (dto.GetCredit, error) {
	// Get user credit by mobile
	c, db := cs.creditRepository.GetByMobile(credit.Mobile)
	if db.Error != nil {
		// Check database error type and handle
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return dto.GetCredit{}, errors.New("404")
		}

		return dto.GetCredit{}, db.Error
	}

	// Assign credit value of credit entity to credit of GetCredit dto
	credit.Credit = c.Credit
	return credit, nil
}

// SetDiscountCredit set user credit by mobile
func (cs *creditService) SetDiscountCredit(credit dto.SetDiscountCredit) (dto.SetDiscountCredit, error) {
	// Set user credit by mobile
	db := cs.creditRepository.UpdateCreditByMobile(credit)
	if db.Error != nil {
		return dto.SetDiscountCredit{}, db.Error
	}

	// check is credit set successfully by mobile or no
	if db.RowsAffected == 1 {
		return credit, nil
	}

	// Create credit if user not exists
	var c entity.Credit = entity.Credit{
		Mobile: credit.Mobile,
		Credit: credit.Credit,
	}

	db = cs.creditRepository.Create(c)
	if db.Error != nil {
		return dto.SetDiscountCredit{}, db.Error
	}

	return credit, nil
}
