package domain_services

import (
	"context"
	"mondu-challenge-alihamedani/domain"
)

type IAccountRepository interface {
	Create(ctx context.Context, account domain.Account) error
	AddTransaction(ctx context.Context, transaction domain.TransactionHistory) error
	GetById(ctx context.Context, id int64) (domain.Account, error)
	GetAll(ctx context.Context) ([]domain.Account, error)
	UpdateBalance(ctx context.Context, account domain.Account) error
}
