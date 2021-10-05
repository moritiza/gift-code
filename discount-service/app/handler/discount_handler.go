package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/moritiza/gift-code/discount-service/app/dto"
	"github.com/moritiza/gift-code/discount-service/app/helper"
	"github.com/moritiza/gift-code/discount-service/app/service"
	"github.com/moritiza/gift-code/discount-service/config"
	"github.com/sirupsen/logrus"
)

type DiscountHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	SubmitDiscount(w http.ResponseWriter, r *http.Request)
}

// discountHandler is a http.Handler and satisfy DiscountHandler interface
type discountHandler struct {
	logger          logrus.Logger
	validator       validator.Validate
	discountService service.DiscountService
}

// NewDiscountHandler creates a new discount handler with the given dependencies
func NewDiscountHandler(l logrus.Logger, v validator.Validate, ds service.DiscountService) DiscountHandler {
	return &discountHandler{
		logger:          l,
		validator:       v,
		discountService: ds,
	}
}

// Create implements the go http.Handler interface
func (dh *discountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var discount dto.Discount

	// Decode received data and store them into Discont DTO
	err := json.NewDecoder(r.Body).Decode(&discount)
	if err != nil {
		// Return 400 error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate received data
	err = dh.validator.Struct(discount)
	if err != nil {
		// Return 400 error with validation errors
		helper.FailureResponse(w, "bad request", config.ValidatorErrors(&dh.validator, err), nil, http.StatusBadRequest)
		return
	}

	// Call discount service Create method
	discount, err = dh.discountService.Create(discount)
	if err != nil {
		// Return 500 error for unhandled errors
		helper.FailureResponse(w, "error", err.Error(), nil, http.StatusInternalServerError)
		return
	}

	// Return Created with header code 201
	helper.SuccessResponse(w, "created", discount, true, http.StatusCreated)
}

// SubmitDiscount implements the go http.Handler interface
func (dh *discountHandler) SubmitDiscount(w http.ResponseWriter, r *http.Request) {
	var discount dto.SubmitDiscount

	// Decode received data and store them into SubmitDiscount DTO
	err := json.NewDecoder(r.Body).Decode(&discount)
	if err != nil {
		// Return 400 error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate received data
	err = dh.validator.Struct(discount)
	if err != nil {
		// Return 400 error with validation errors
		helper.FailureResponse(w, "bad request", config.ValidatorErrors(&dh.validator, err), nil, http.StatusBadRequest)
		return
	}

	// Call discount service SubmitDiscount method
	discount, err = dh.discountService.SubmitDiscount(discount)
	if err != nil {
		switch err.Error() {
		case "404":
			// Return 404 error
			helper.FailureResponse(w, "not found", "discount not found", nil, http.StatusNotFound)
			return
		case "finished":
			// Return 400 error
			helper.FailureResponse(w, "bad request", "discount code has been finished", nil, http.StatusBadRequest)
			return
		case "used":
			// Return 400 error
			helper.FailureResponse(w, "bad request", "discount code has been used", nil, http.StatusBadRequest)
			return
		default:
			// Return 500 error for unhandled errors
			helper.FailureResponse(w, "error", err.Error(), nil, http.StatusInternalServerError)
			return
		}
	}

	// Return OK with header code 200
	helper.SuccessResponse(w, "ok", discount, true, http.StatusOK)
}
