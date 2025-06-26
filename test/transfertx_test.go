package test

import (
	"context"
	"testing"

	"github.com/leoneIAguilera/simple_bank/internal/db"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := db.NewStore(testDB)
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	amount := int64(10)
	n := int64(5)

	errs := make(chan error)
	results := make(chan db.TransferTxResults)

	for range n {
		go func() {
			result, err := store.TransferTx(context.Background(), db.TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}

	for range n {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		//check transfer
		require.Equal(t, account1.ID, result.Transfer.FromAccountID)
		require.Equal(t, account2.ID, result.Transfer.ToAccountID)
		require.Equal(t, amount, result.Transfer.Amount)
		require.NotZero(t, result.Transfer.ID)
		require.NotZero(t, result.Transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), result.Transfer.ID)
		require.NoError(t, err)

		// check entries

	}
}
