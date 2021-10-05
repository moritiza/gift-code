package service

import (
	"github.com/moritiza/gift-code/report-service/app/dto"
	"github.com/moritiza/gift-code/report-service/app/entity"
	"github.com/moritiza/gift-code/report-service/app/repository"
	"github.com/sirupsen/logrus"
)

type ReportService interface {
	Get(page int, size int) (dto.Reports, error)
	Create(report dto.Report) (dto.Report, error)
}

// reportService satisfy ReportService interface
type reportService struct {
	logger           logrus.Logger
	reportRepository repository.ReportRepository
}

// NewReportService creates a new report service with the given dependencies
func NewReportService(l logrus.Logger, rr repository.ReportRepository) ReportService {
	return &reportService{
		logger:           l,
		reportRepository: rr,
	}
}

// Get return all reports
func (rs *reportService) Get(page int, size int) (dto.Reports, error) {
	// Get all reports
	reports, db := rs.reportRepository.Get(page, size)
	if db.Error != nil {
		return dto.Reports{}, db.Error
	}

	return dto.Reports{Reports: reports}, nil
}

// Create do creating report
func (rs *reportService) Create(report dto.Report) (dto.Report, error) {
	var re = entity.Report{
		Mobile:     report.Mobile,
		Code:       report.Code,
		CodeCredit: report.CodeCredit,
	}

	// Insert new report to reports table
	db := rs.reportRepository.Create(re)
	if db.Error != nil {
		rs.logger.Error("Error: ", db.Error)
		return dto.Report{}, db.Error
	}

	return report, nil
}
