package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"mondu-challenge-alihamedani/application/contracts"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	arg := contracts.CreateAccountModel{
		FirstName: "x",
		LastName:  "y",
		IBAN:      "z",
		Balance:   10,
	}

	account, err := accountService.AddAccount(context.Background(), &arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.FirstName, account.FirstName)
	require.Equal(t, arg.LastName, account.LastName)
	require.Equal(t, arg.IBAN, account.IBAN)
	require.Equal(t, arg.Balance, account.Balance)

	require.NotZero(t, account.Id)
	require.NotZero(t, account.CreationDate)
}
