package core

import (
	"github.com/moritiza/gift-code/report-service/app/handler"
	"github.com/moritiza/gift-code/report-service/app/repository"
	"github.com/moritiza/gift-code/report-service/app/service"
	"github.com/moritiza/gift-code/report-service/config"
)

// Dependencies store all dependencies
type Dependencies struct {
	Repositories Repositories
	Services     Services
	Handlers     Handlers
}

// Repositories store all repositories
type Repositories struct {
	ReportRepository repository.ReportRepository
}

// Services store all services
type Services struct {
	ReportService service.ReportService
}

// Handlers store all handlers
type Handlers struct {
	ReportHandler handler.ReportHandler
}

// PrepareDependensies prepare application necessary dependencies
func PrepareDependensies(cfg config.Config) *Dependencies {
	var (
		repositories Repositories
		services     Services
		handlers     Handlers
	)

	repositories.ReportRepository = repository.NewReportRepository(cfg.Database)
	services.ReportService = service.NewReportService(*cfg.Logger, repositories.ReportRepository)
	handlers.ReportHandler = handler.NewReportHandler(*cfg.Logger, *cfg.Validator, services.ReportService)

	return &Dependencies{
		Repositories: repositories,
		Services:     services,
		Handlers:     handlers,
	}
}
