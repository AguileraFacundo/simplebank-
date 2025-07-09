package test

import (
	"context"
	"testing"

	"github.com/leoneIAguilera/simple_bank/internal/db"
	"github.com/leoneIAguilera/simple_bank/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) db.Transfer {
	arg := db.CreateTransferParams{
		FromAccountID: util.RandomNumber(1, 5),
		ToAccountID:   util.RandomNumber(1, 5),
		Amount:        util.RandomMoney(),
	}

	transfer, err := TestQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)
	check, err := TestQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.Equal(t, transfer.ID, check.ID)
	require.Equal(t, transfer.FromAccountID, check.FromAccountID)
	require.Equal(t, transfer.ToAccountID, check.ToAccountID)
	require.Equal(t, transfer.Amount, check.Amount)
}

func TestListTransfer(t *testing.T) {
	for range 10 {
		createRandomTransfer(t)
	}

	arg := db.ListTransferParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := TestQueries.ListTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	err := TestQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	check, err := TestQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.Empty(t, check)
}
