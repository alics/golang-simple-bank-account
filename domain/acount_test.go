package domain

import "testing"

func TestAccount_Deposit(t *testing.T) {
	t.Parallel()

	type args struct {
		amount float64
	}

	tests := []struct {
		name     string
		account  Account
		args     args
		expected float64
	}{
		{
			name: "successful depositing balance",
			args: args{
				amount: 10,
			},
			account:  SetAccountBalance(0),
			expected: 10,
		},
		{
			name: "successful depositing balance",
			args: args{
				amount: 102098,
			},
			account:  SetAccountBalance(0),
			expected: 102098,
		},
		{
			name: "successful depositing balance",
			args: args{
				amount: 4498,
			},
			account:  SetAccountBalance(98),
			expected: 4596,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.account.Deposit(tt.args.amount)

			if tt.account.Balance() != tt.expected {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					tt.account.Balance(),
					tt.expected,
				)
			}
		})
	}
}

func TestAccount_Withdraw(t *testing.T) {
	t.Parallel()

	type args struct {
		amount float64
	}

	tests := []struct {
		name        string
		account     Account
		args        args
		expected    float64
		expectedErr error
	}{
		{
			name: "success in withdrawing balance",
			args: args{
				amount: 10,
			},
			account:  SetAccountBalance(10),
			expected: 0,
		},
		{
			name: "success in withdrawing balance",
			args: args{
				amount: 10012,
			},
			account:  SetAccountBalance(10013),
			expected: 1,
		},
		{
			name: "success in withdrawing balance",
			args: args{
				amount: 25,
			},
			account:  SetAccountBalance(125),
			expected: 100,
		},
		{
			name: "error when withdrawing account balance without sufficient balance",
			args: args{
				amount: 564,
			},
			account:     SetAccountBalance(62),
			expectedErr: InsufficientBalanceError,
		},
		{
			name: "error when withdrawing account balance without sufficient balance",
			args: args{
				amount: 5,
			},
			account:     SetAccountBalance(1),
			expectedErr: InsufficientBalanceError,
		},
		{
			name: "error when withdrawing account balance without sufficient balance",
			args: args{
				amount: 10,
			},
			account:     SetAccountBalance(0),
			expectedErr: InsufficientBalanceError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if err = tt.account.Withdraw(tt.args.amount); (err != nil) && (err.Error() != tt.expectedErr.Error()) {
				t.Errorf("[TestCase '%s'] ResultError: '%v' | ExpectedError: '%v'",
					tt.name,
					err,
					tt.expectedErr.Error(),
				)
				return
			}

			if tt.expectedErr == nil && tt.account.Balance() != tt.expected {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					tt.account.Balance(),
					tt.expected,
				)
			}
		})
	}
}
