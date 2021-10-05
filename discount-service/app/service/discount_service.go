package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/moritiza/gift-code/discount-service/app/dto"
	"github.com/moritiza/gift-code/discount-service/app/entity"
	"github.com/moritiza/gift-code/discount-service/app/helper"
	"github.com/moritiza/gift-code/discount-service/app/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DiscountService interface {
	Create(report dto.Discount) (dto.Discount, error)
	SubmitDiscount(discount dto.SubmitDiscount) (dto.SubmitDiscount, error)
}

// discountService satisfy DiscountService interface
type discountService struct {
	logger                 logrus.Logger
	discountRepository     repository.DiscountRepository
	usedDiscountRepository repository.UsedDiscountRepository
}

// NewDiscountService creates a new discount service with the given dependencies
func NewDiscountService(l logrus.Logger, dr repository.DiscountRepository, udr repository.UsedDiscountRepository) DiscountService {
	return &discountService{
		logger:                 l,
		discountRepository:     dr,
		usedDiscountRepository: udr,
	}
}

// Create do creating discount
func (ds *discountService) Create(discount dto.Discount) (dto.Discount, error) {
	var de = entity.Discount{
		Code:       discount.Code,
		CodeCredit: discount.CodeCredit,
		Count:      discount.Count,
	}

	// Insert new discount to discounts table
	db := ds.discountRepository.Create(de)
	if db.Error != nil {
		ds.logger.Error("Error: ", db.Error)
		return dto.Discount{}, db.Error
	}

	return discount, nil
}

// SubmitDiscount do submitting discount steps
func (ds *discountService) SubmitDiscount(discount dto.SubmitDiscount) (dto.SubmitDiscount, error) {
	// Check is discount code used by user or no?
	// If database error is not found error -> user can use of the code
	_, db := ds.usedDiscountRepository.GetByCodeAndMobile(discount)
	if db.Error != nil {
		// Check database error type and handle
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			d, err := ds.getDiscountByCode(discount.Code)
			if err != nil {
				return dto.SubmitDiscount{}, err
			}

			err = ds.checkDiscountCount(d)
			if err != nil {
				return dto.SubmitDiscount{}, err
			}

			err = ds.setDiscountCreditForUser(d, discount)
			if err != nil {
				return dto.SubmitDiscount{}, err
			}

			err = ds.createReport(d, discount)
			if err != nil {
				return dto.SubmitDiscount{}, err
			}

			err = ds.incrementCodeUsed(d)
			if err != nil {
				return dto.SubmitDiscount{}, err
			}

			err = ds.createUsedDiscount(discount)
			if err != nil {
				return dto.SubmitDiscount{}, err
			}

			return discount, nil
		}

		return dto.SubmitDiscount{}, db.Error
	}

	return dto.SubmitDiscount{}, errors.New("used")
}

// getDiscountByCode is here to find discount by discount code
func (ds *discountService) getDiscountByCode(code string) (entity.Discount, error) {
	d, db := ds.discountRepository.GetByCode(code)
	if db.Error != nil {
		// Check database error type and handle
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return entity.Discount{}, errors.New("404")
		}

		return entity.Discount{}, db.Error
	}

	return d, nil
}

// checkDiscountCount is here for check is discount code already finished or not
func (ds *discountService) checkDiscountCount(discount entity.Discount) error {
	if discount.Count > discount.Used {
		return nil
	}
	return errors.New("finished")
}

// setDiscountCreditForUser send request to wallet service to set discount code credit to user
func (ds *discountService) setDiscountCreditForUser(d entity.Discount, discount dto.SubmitDiscount) error {
	// Prepare credit data
	credit := dto.SetDiscountCredit{
		Mobile: discount.Mobile,
		Credit: d.CodeCredit,
	}

	// Create json request data
	reqBody, err := json.Marshal(credit)
	if err != nil {
		return err
	}

	// Send request to report service to create a new report
	resp, err := helper.Rester(
		http.MethodPatch,
		os.Getenv("WALLET_HOST")+":"+os.Getenv("WALLET_PORT")+"/api/credit/set-discount",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check is user credit set successfully or not
	if resp.StatusCode == 200 {
		return nil
	}
	return errors.New("wallet")
}

// createReport create report on report service
func (ds *discountService) createReport(d entity.Discount, discount dto.SubmitDiscount) error {
	// Prepare report data
	report := dto.Report{
		Mobile:     discount.Mobile,
		Code:       d.Code,
		CodeCredit: d.CodeCredit,
	}

	// Create json request data
	reqBody, err := json.Marshal(report)
	if err != nil {
		return err
	}

	// Send request to report service to create a new report
	resp, err := helper.Rester(
		http.MethodPost,
		os.Getenv("REPORT_HOST")+":"+os.Getenv("REPORT_PORT")+"/api/report",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check is report created successfully or not
	if resp.StatusCode == 201 {
		return nil
	}
	return errors.New("report")
}

// incrementCodeUsed is here for increment discount code used one unit
func (ds *discountService) incrementCodeUsed(discount entity.Discount) error {
	db := ds.discountRepository.IncrementCodeUsed(discount.ID)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

// createUsedDiscount create used discount to handle that user not reuse the code
func (ds *discountService) createUsedDiscount(discount dto.SubmitDiscount) error {
	var ud entity.UsedDiscount = entity.UsedDiscount{Code: discount.Code, Mobile: discount.Mobile}

	db := ds.usedDiscountRepository.Create(ud)
	if db.Error != nil {
		return db.Error
	}

	return nil
}
