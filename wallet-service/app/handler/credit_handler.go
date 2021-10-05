package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/moritiza/gift-code/wallet-service/app/dto"
	"github.com/moritiza/gift-code/wallet-service/app/helper"
	"github.com/moritiza/gift-code/wallet-service/app/service"
	"github.com/moritiza/gift-code/wallet-service/config"
	"github.com/sirupsen/logrus"
)

type CreditHandler interface {
	GetCredit(w http.ResponseWriter, r *http.Request)
	SetDiscountCredit(w http.ResponseWriter, r *http.Request)
}

// creditHandler is a http.Handler and satisfy CreditHandler interface
type creditHandler struct {
	logger        logrus.Logger
	validator     validator.Validate
	creditService service.CreditService
}

// NewCreditHandler creates a new credit handler with the given dependencies
func NewCreditHandler(l logrus.Logger, v validator.Validate, cs service.CreditService) CreditHandler {
	return &creditHandler{
		logger:        l,
		validator:     v,
		creditService: cs,
	}
}

// GetCredit implements the go http.Handler interface
func (ch *creditHandler) GetCredit(w http.ResponseWriter, r *http.Request) {
	var credit dto.GetCredit

	// Decode received data and store them into GetCredit DTO
	err := json.NewDecoder(r.Body).Decode(&credit)
	if err != nil {
		// Return 400 error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate received data
	err = ch.validator.Struct(credit)
	if err != nil {
		// Return 400 error with validation errors
		helper.FailureResponse(w, "bad request", config.ValidatorErrors(&ch.validator, err), nil, http.StatusBadRequest)
		return
	}

	// Call credit service GetCredit method
	credit, err = ch.creditService.GetCredit(credit)
	if err != nil {
		switch err.Error() {
		case "404":
			// Return 404 error
			helper.FailureResponse(w, "not found", "user not found", nil, http.StatusNotFound)
			return
		default:
			// Return 500 error for unhandled errors
			helper.FailureResponse(w, "error", err.Error(), nil, http.StatusInternalServerError)
			return
		}
	}

	// Return OK with header code 200
	helper.SuccessResponse(w, "ok", credit, true, http.StatusOK)
}

// SetDiscountCredit implements the go http.Handler interface
func (ch *creditHandler) SetDiscountCredit(w http.ResponseWriter, r *http.Request) {
	var credit dto.SetDiscountCredit

	// Decode received data and store them into SetDiscountCredit DTO
	err := json.NewDecoder(r.Body).Decode(&credit)
	if err != nil {
		// Return 400 error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate received data
	err = ch.validator.Struct(credit)
	if err != nil {
		// Return 400 error with validation errors
		helper.FailureResponse(w, "bad request", config.ValidatorErrors(&ch.validator, err), nil, http.StatusBadRequest)
		return
	}

	// Call credit service SetDiscountCredit method
	credit, err = ch.creditService.SetDiscountCredit(credit)
	if err != nil {
		// Return 500 error for unhandled errors
		helper.FailureResponse(w, "error", err.Error(), nil, http.StatusInternalServerError)
		return
	}

	// Return OK with header code 200
	helper.SuccessResponse(w, "ok", credit, true, http.StatusOK)
}
