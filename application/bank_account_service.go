package application

import (
	"context"
	"mondu-challenge-alihamedani/application/contracts"
	"mondu-challenge-alihamedani/domain"
	"mondu-challenge-alihamedani/domain/domain_services"
	util "mondu-challenge-alihamedani/infrastructure/utils"
	"time"
)

const (
	Withdraw int16 = 0
	Deposit  int16 = 1
)

type (
	BankAccountService interface {
		AddAccount(ctx context.Context, model *contracts.CreateAccountModel) (contracts.AccountResult, error)
		DepositMoneyToAccount(ctx context.Context, model *contracts.CreateTransferModel) error
		WithdrawMoneyFromAccount(ctx context.Context, model *contracts.CreateTransferModel) error
		GetBalance(ctx context.Context, model *contracts.GetAccountBalanceModel) (contracts.AccountBalanceResult, error)
		GetAllAccounts(ctx context.Context) ([]contracts.AccountResult, error)
	}
	bankAccountService struct {
		Repository domain_services.IAccountRepository
	}
)

func NewAccountService(repository domain_services.IAccountRepository) BankAccountService {
	return &bankAccountService{Repository: repository}
}

func (b bankAccountService) GetAllAccounts(ctx context.Context) ([]contracts.AccountResult, error) {
	var result []contracts.AccountResult

	accountList, err := b.Repository.GetAll(ctx)
	if err != nil {
		return result, err
	}

	for _, item := range accountList {
		account := contracts.AccountResult{
			Id:           item.ID(),
			FirstName:    item.FirstName(),
			LastName:     item.LastName(),
			IBAN:         item.IBAN(),
			Balance:      item.Balance(),
			CreationDate: item.CreationDate(),
		}
		result = append(result, account)
	}
	return result, nil
}

func (b bankAccountService) GetBalance(ctx context.Context, model *contracts.GetAccountBalanceModel) (contracts.AccountBalanceResult, error) {
	account, err := b.Repository.GetById(ctx, model.Id)

	if err != nil {
		return contracts.AccountBalanceResult{}, err
	}
	if account.ID() == 0 {
		return contracts.AccountBalanceResult{}, domain.AccountNotFoundError
	}

	result := contracts.AccountBalanceResult{
		FirstName: account.FirstName(),
		LastName:  account.LastName(),
		IBAN:      account.IBAN(),
		Balance:   account.Balance(),
	}
	return result, err
}

func (b bankAccountService) WithdrawMoneyFromAccount(ctx context.Context, model *contracts.CreateTransferModel) error {
	account, err := b.Repository.GetById(ctx, model.SourceAccountId)

	if err != nil {
		return err
	}
	if account.ID() == 0 {
		return domain.AccountNotFoundError
	}

	account.Withdraw(model.Amount)
	err = b.Repository.UpdateBalance(ctx, account)

	go func() {
		var transaction = domain.AddTransaction(util.NewId(), model.SourceAccountId, model.DestinationAccountId, model.Amount, account.Balance(), Withdraw)
		err = b.Repository.AddTransaction(ctx, transaction)
	}()

	return err
}

func (b bankAccountService) DepositMoneyToAccount(ctx context.Context, model *contracts.CreateTransferModel) error {
	account, err := b.Repository.GetById(ctx, model.SourceAccountId)
	if err != nil {
		return err
	}
	if account.ID() == 0 {
		return domain.AccountNotFoundError
	}

	account.Deposit(model.Amount)
	err = b.Repository.UpdateBalance(ctx, account)

	go func() {
		var transaction = domain.AddTransaction(util.NewId(), model.SourceAccountId, model.DestinationAccountId, model.Amount, account.Balance(), Deposit)
		b.Repository.AddTransaction(ctx, transaction)
	}()

	return err
}

func (b bankAccountService) AddAccount(ctx context.Context, model *contracts.CreateAccountModel) (contracts.AccountResult, error) {
	id := util.NewId()
	account := domain.CreateAccount(id, model.FirstName, model.LastName, model.IBAN, model.Balance, time.Now())
	err := b.Repository.Create(ctx, account)

	result := contracts.AccountResult{
		Id:           account.ID(),
		FirstName:    account.FirstName(),
		LastName:     account.LastName(),
		IBAN:         account.IBAN(),
		Balance:      account.Balance(),
		CreationDate: account.CreationDate(),
	}
	return result, err
}
