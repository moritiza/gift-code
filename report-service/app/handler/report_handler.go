package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/moritiza/gift-code/report-service/app/dto"
	"github.com/moritiza/gift-code/report-service/app/helper"
	"github.com/moritiza/gift-code/report-service/app/service"
	"github.com/moritiza/gift-code/report-service/config"
	"github.com/sirupsen/logrus"
)

type ReportHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

// reportHandler is a http.Handler and satisfy ReportHandler interface
type reportHandler struct {
	logger        logrus.Logger
	validator     validator.Validate
	reportService service.ReportService
}

// NewReportHandler creates a new report handler with the given dependencies
func NewReportHandler(l logrus.Logger, v validator.Validate, rs service.ReportService) ReportHandler {
	return &reportHandler{
		logger:        l,
		validator:     v,
		reportService: rs,
	}
}

// GetCredit implements the go http.Handler interface
func (rh *reportHandler) Get(w http.ResponseWriter, r *http.Request) {
	var (
		page int = 1
		size int = 10
		err  error
	)

	// Get pagination page number
	p := r.FormValue("page")
	if p != "" {
		page, err = strconv.Atoi(p)
		if err != nil {
			helper.FailureResponse(w, "error", err.Error(), nil, http.StatusInternalServerError)
			return
		}
		if page == 0 {
			page = 1
		}
	}

	// Get pagination page size
	s := r.FormValue("size")
	if s != "" {
		size, err = strconv.Atoi(s)
		if err != nil {
			helper.FailureResponse(w, "error", err.Error(), nil, http.StatusInternalServerError)
			return
		}
		if size == 0 {
			size = 10
		}
	}

	// Call report service Get method
	reports, err := rh.reportService.Get(page, size)
	if err != nil {
		rh.logger.Debug("here4")
		rh.logger.Debug("Error: ", err)
		helper.FailureResponse(w, "error", err.Error(), nil, http.StatusInternalServerError)
		return
	}

	// Return OK with header code 200
	helper.SuccessResponse(w, "ok", reports, true, http.StatusOK)
}

// Create implements the go http.Handler interface
func (rh *reportHandler) Create(w http.ResponseWriter, r *http.Request) {
	var report dto.Report

	// Decode received data and store them into Report DTO
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		// Return 400 error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate received data
	err = rh.validator.Struct(report)
	if err != nil {
		// Return 400 error with validation errors
		helper.FailureResponse(w, "bad request", config.ValidatorErrors(&rh.validator, err), nil, http.StatusBadRequest)
		return
	}

	// Call report service Create method
	report, err = rh.reportService.Create(report)
	if err != nil {
		// Return 500 error for unhandled errors
		helper.FailureResponse(w, "error", err.Error(), nil, http.StatusInternalServerError)
		return
	}

	// Return Created with header code 201
	helper.SuccessResponse(w, "created", report, true, http.StatusCreated)
}
