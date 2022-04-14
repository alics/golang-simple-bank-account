package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"mondu-challenge-alihamedani/domain"
	"mondu-challenge-alihamedani/domain/domain_services"
	"time"
)

type accountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) domain_services.IAccountRepository {
	return &accountRepository{DB: db}
}

func (r accountRepository) GetAll(ctx context.Context) ([]domain.Account, error) {
	var accountList []domain.Account

	query := fmt.Sprintf("SELECT Id,FirstName,LastName,IBAN,Balance,CreationDate FROM [BankAccount].[dbo].[Accounts] ")
	rows, err := r.DB.Query(query)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return accountList, err
	}

	defer rows.Close()

	for rows.Next() {
		var Id int64
		var FirstName, LastName, IBAN string
		var Balance float64
		var CreationDate time.Time
		err := rows.Scan(&Id, &FirstName, &LastName, &IBAN, &Balance, &CreationDate)
		if err != nil {
			log.Fatal("Error reading rows: " + err.Error())
			return accountList, err
		}
		account := domain.CreateAccount(Id, FirstName, LastName, IBAN, Balance, CreationDate)
		accountList = append(accountList, account)
	}
	return accountList, nil
}

func (r accountRepository) Create(ctx context.Context, account domain.Account) error {
	command := fmt.Sprintf("INSERT INTO [Accounts]([Id],[FirstName],[LastName],[IBAN],[Balance],[CreationDate])"+
		" VALUES ('%d','%s','%s','%s',%f,'%s')",
		account.ID(), account.FirstName(), account.LastName(), account.IBAN(), account.Balance(), account.CreationDate().Format(time.RFC3339))

	_, err := r.DB.ExecContext(ctx, command)
	if err != nil {
		return err
	}

	return nil
}

func (r accountRepository) AddTransaction(ctx context.Context, transaction domain.TransactionHistory) error {
	command := fmt.Sprintf("INSERT INTO [TransactionHistories]([Id],[SourceAccountId],[DestinationAccountId],[TransactionType],[Amount],[CurrentBalance],[TransactionDateTime])"+
		" VALUES (%d,%d,%d,%d,%f,%f,'%s')",
		transaction.ID(), transaction.SourceAccountId(), transaction.DestinationAccountId(), transaction.TransactionType(), transaction.Amount(), transaction.CurrentBalance(), transaction.TransactionDateTime().Format(time.RFC3339))

	_, err := r.DB.ExecContext(ctx, command)
	if err != nil {
		return err
	}

	return nil
}

func (r accountRepository) GetById(ctx context.Context, id int64) (domain.Account, error) {
	account := domain.Account{}

	query := fmt.Sprintf("SELECT Id,FirstName,LastName,IBAN,Balance,CreationDate FROM [BankAccount].[dbo].[Accounts] WHERE Id='%d'", id)
	rows, err := r.DB.Query(query)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return account, err
	}

	defer rows.Close()

	for rows.Next() {
		var Id int64
		var FirstName, LastName, IBAN string
		var Balance float64
		var CreationDate time.Time
		err := rows.Scan(&Id, &FirstName, &LastName, &IBAN, &Balance, &CreationDate)
		if err != nil {
			log.Fatal("Error reading rows: " + err.Error())
			return account, err
		}
		account = domain.CreateAccount(Id, FirstName, LastName, IBAN, Balance, CreationDate)
		return account, nil
	}
	return account, nil
}

func (r accountRepository) UpdateBalance(ctx context.Context, account domain.Account) error {
	command := fmt.Sprintf(" UPDATE BankAccount.dbo.Accounts SET Balance=%f WHERE Id=%d",
		account.Balance(), account.ID())

	_, err := r.DB.Exec(command)
	if err != nil {
		log.Fatal("Error updating row: " + err.Error())
		return err
	}
	return nil
}
