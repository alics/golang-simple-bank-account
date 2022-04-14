package domain

import "time"

type TransactionId string

func (t TransactionId) String() string {
	return string(t)
}

type TransactionHistory struct {
	id                   int64
	sourceAccountId      int64
	destinationAccountId int64
	transactionType      int16
	amount               float64
	currentBalance       float64
	transactionDateTime  time.Time
}

func AddTransaction(id, sourceAccountId, destinationAccountId int64, amount float64, currentBalance float64, transactionType int16) TransactionHistory {
	return TransactionHistory{
		id:                   id,
		sourceAccountId:      sourceAccountId,
		destinationAccountId: destinationAccountId,
		amount:               amount,
		currentBalance:       currentBalance,
		transactionType:      transactionType,
		transactionDateTime:  time.Now(),
	}
}

func (t TransactionHistory) ID() int64 {
	return t.id
}

func (t TransactionHistory) SourceAccountId() int64 {
	return t.sourceAccountId
}

func (t TransactionHistory) DestinationAccountId() int64 {
	return t.destinationAccountId
}

func (t TransactionHistory) TransactionType() int16 {
	return t.transactionType
}

func (t TransactionHistory) Amount() float64 {
	return t.amount
}
func (t TransactionHistory) CurrentBalance() float64 {
	return t.currentBalance
}

func (t TransactionHistory) TransactionDateTime() time.Time {
	return t.transactionDateTime
}
