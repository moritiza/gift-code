package core

import (
	"github.com/moritiza/gift-code/discount-service/app/handler"
	"github.com/moritiza/gift-code/discount-service/app/repository"
	"github.com/moritiza/gift-code/discount-service/app/service"
	"github.com/moritiza/gift-code/discount-service/config"
)

// Dependencies store all dependencies
type Dependencies struct {
	Repositories Repositories
	Services     Services
	Handlers     Handlers
}

// Repositories store all repositories
type Repositories struct {
	discountRepository     repository.DiscountRepository
	usedDiscountRepository repository.UsedDiscountRepository
}

// Services store all services
type Services struct {
	discountService service.DiscountService
}

// Handlers store all handlers
type Handlers struct {
	discountHandler handler.DiscountHandler
}

// PrepareDependensies prepare application necessary dependencies
func PrepareDependensies(cfg config.Config) *Dependencies {
	var (
		repositories Repositories
		services     Services
		handlers     Handlers
	)

	repositories.discountRepository = repository.NewDiscountRepository(cfg.Database)
	repositories.usedDiscountRepository = repository.NewUsedDiscountRepository(cfg.Database)
	services.discountService = service.NewDiscountService(*cfg.Logger, repositories.discountRepository, repositories.usedDiscountRepository)
	handlers.discountHandler = handler.NewDiscountHandler(*cfg.Logger, *cfg.Validator, services.discountService)

	return &Dependencies{
		Repositories: repositories,
		Services:     services,
		Handlers:     handlers,
	}
}
