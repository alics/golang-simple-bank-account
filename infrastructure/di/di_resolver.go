package di

import (
	"mondu-challenge-alihamedani/application"
	"mondu-challenge-alihamedani/domain/domain_services"
	"mondu-challenge-alihamedani/infrastructure/persistence"
	util "mondu-challenge-alihamedani/infrastructure/utils"
)

func NewAccountRepositoryResolve(config util.Config) domain_services.IAccountRepository {
	return persistence.NewAccountRepository(persistence.InitSqlDB(config))
}

func NewAccountServiceResolve(config util.Config) application.BankAccountService {
	return application.NewAccountService(NewAccountRepositoryResolve(config))
}
