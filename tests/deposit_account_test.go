package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"mondu-challenge-alihamedani/application/contracts"
	"testing"
)

func TestDepositAccount(t *testing.T) {
	allAccounts1, _ := accountService.GetAllAccounts(context.Background())
	forTestAccount := allAccounts1[0]

	arg := contracts.CreateTransferModel{
		SourceAccountId:      forTestAccount.Id,
		DestinationAccountId: 123456789,
		Amount:               1,
	}

	err := accountService.DepositMoneyToAccount(context.Background(), &arg)
	require.NoError(t, err)

	model := contracts.GetAccountBalanceModel{Id: forTestAccount.Id}
	balance, _ := accountService.GetBalance(context.Background(), &model)

	require.Equal(t, forTestAccount.Balance+arg.Amount, balance.Balance)
}
