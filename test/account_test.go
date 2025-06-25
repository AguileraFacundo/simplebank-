package test

import (
	"context"
	"testing"
	"time"

	"github.com/leoneIAguilera/simple_bank/internal/db"
	"github.com/leoneIAguilera/simple_bank/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) db.Account {
	arg := db.CreateAccountParams{
		Owner:    util.CreateRandomOwner(),
		Balance:  util.CreateRandomMoney(),
		Currency: util.CreateRandomCurrencies(),
	}

	account, err := TestQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	for range 5 {
		createRandomAccount(t)
	}
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	account1, err := TestQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	require.Equal(t, account.ID, account1.ID)
	require.Equal(t, account.Owner, account1.Owner)
	require.Equal(t, account.Balance, account1.Balance)
	require.Equal(t, account.Currency, account1.Currency)
	require.WithinDuration(t, account.CreatedAt, account1.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T) {
	for range 5 {
		createRandomAccount(t)
	}

	arg := db.ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := TestQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := TestQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account1, err := TestQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, account1)
	require.NotEqual(t, account.ID, account1.ID)
	require.NotEqual(t, account.Owner, account1.Owner)
	require.NotEqual(t, account.Balance, account1.Balance)
	require.NotEqual(t, account.Currency, account1.Currency)

}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := db.UpdateAccountParams{
		ID:      account.ID,
		Balance: 300,
	}

	account1, err := TestQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEqual(t, account.Balance, account1.Balance)
}
