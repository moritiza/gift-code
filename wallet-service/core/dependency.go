package core

import (
	"github.com/moritiza/gift-code/wallet-service/app/handler"
	"github.com/moritiza/gift-code/wallet-service/app/repository"
	"github.com/moritiza/gift-code/wallet-service/app/service"
	"github.com/moritiza/gift-code/wallet-service/config"
)

// Dependencies store all dependencies
type Dependencies struct {
	Repositories Repositories
	Services     Services
	Handlers     Handlers
}

// Repositories store all repositories
type Repositories struct {
	CreditRepository repository.CreditRepository
}

// Services store all services
type Services struct {
	Creditservice service.CreditService
}

// Handlers store all handlers
type Handlers struct {
	CreditHandler handler.CreditHandler
}

// PrepareDependensies prepare application necessary dependencies
func PrepareDependensies(cfg config.Config) *Dependencies {
	var (
		repositories Repositories
		services     Services
		handlers     Handlers
	)

	repositories.CreditRepository = repository.NewCreditRepository(cfg.Database)
	services.Creditservice = service.NewCreditService(*cfg.Logger, repositories.CreditRepository)
	handlers.CreditHandler = handler.NewCreditHandler(*cfg.Logger, *cfg.Validator, services.Creditservice)

	return &Dependencies{
		Repositories: repositories,
		Services:     services,
		Handlers:     handlers,
	}
}
