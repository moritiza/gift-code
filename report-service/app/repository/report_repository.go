package repository

import (
	"github.com/moritiza/gift-code/report-service/app/dto"
	"github.com/moritiza/gift-code/report-service/app/entity"
	"gorm.io/gorm"
)

type ReportRepository interface {
	Get(page int, size int) ([]dto.Report, *gorm.DB)
	Create(report entity.Report) *gorm.DB
}

// reportRepository satisfy ReportRepository interface
type reportRepository struct {
	db *gorm.DB
}

// NewReportRepository creates a new report repository with the given dependencies
func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{
		db: db,
	}
}

// Get do read operation on reports table and return all reports with database result
func (rr *reportRepository) Get(page int, size int) ([]dto.Report, *gorm.DB) {
	var report []dto.Report

	r := rr.db.Model(entity.Report{}).Order("id DESC").Offset((page - 1) * size).Limit(size).Find(&report)
	return report, r
}

// Create do insert operation on reports table and return database result
func (rr *reportRepository) Create(report entity.Report) *gorm.DB {
	r := rr.db.Model(entity.Report{}).Create(&report)
	return r
}
