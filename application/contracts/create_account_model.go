package contracts

import "time"

type AccountResult struct {
	Id           int64     `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	IBAN         string    `json:"iban"`
	Balance      float64   `json:"balance"`
	CreationDate time.Time `json:"creation_date"`
}

type AccountBalanceResult struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	IBAN      string  `json:"iban"`
	Balance   float64 `json:"balance"`
}

type GetAccountBalanceModel struct {
	Id int64 `json:"id"`
}

type CreateAccountModel struct {
	FirstName string  `json:"first_name" validate:"required"`
	LastName  string  `json:"last_name" validate:"required"`
	IBAN      string  `json:"iban" validate:"required"`
	Balance   float64 `json:"balance" validate:"gt=0,required"`
}

type CreateTransferModel struct {
	SourceAccountId      int64   `json:"source_account_id" validate:"required"`
	DestinationAccountId int64   `json:"destination_account_id" validate:"required"`
	Amount               float64 `json:"amount" validate:"gt=0,required"`
}
