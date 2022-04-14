package domain

import (
	"errors"
	"time"
)

var (
	AccountNotFoundError            = errors.New("account not found")
	SourceAccountNotFoundError      = errors.New("source account not found")
	DestinationAccountNotFoundError = errors.New("destination account not found")
	InsufficientBalanceError        = errors.New("source account does not have sufficient balance")
)

type Account struct {
	id           int64
	firstName    string
	lastName     string
	iban         string
	balance      float64
	creationDate time.Time
}

type AccountId string

func (a AccountId) String() string {
	return string(a)
}

func CreateAccount(id int64, firstName, lastName, iban string, balance float64, creationDate time.Time) Account {
	return Account{
		id:           id,
		firstName:    firstName,
		lastName:     lastName,
		iban:         iban,
		balance:      balance,
		creationDate: creationDate,
	}
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) error {
	if a.balance < amount {
		return InsufficientBalanceError
	}
	a.balance -= amount
	return nil
}

func (a Account) ID() int64 {
	return a.id
}

func (a Account) FirstName() string {
	return a.firstName
}

func (a Account) LastName() string {
	return a.lastName
}

func (a Account) IBAN() string {
	return a.iban
}

func (a Account) Balance() float64 {
	return a.balance
}

func (a Account) CreationDate() time.Time {
	return a.creationDate
}

func SetAccountBalance(balance float64) Account {
	return Account{balance: balance}
}
