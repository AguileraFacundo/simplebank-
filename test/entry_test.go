package test

import (
	"context"
	"testing"

	"github.com/leoneIAguilera/simple_bank/internal/db"
	"github.com/leoneIAguilera/simple_bank/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) db.Entry {
	arg := db.CreateEntryParams{
		AccountID: util.CreateRandomNumber(1, 20),
		Amount:    util.CreateRandomMoney(),
	}
	entry, err := TestQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, entry.ID)
	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestDeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)
	err := TestQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	check, err := TestQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.Empty(t, check)
}

func TestGetEntry(t *testing.T) {
	entry := createRandomEntry(t)
	check, err := TestQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.Equal(t, entry.ID, check.ID)
	require.Equal(t, entry.AccountID, check.AccountID)
	require.Equal(t, entry.Amount, check.Amount)
	require.Equal(t, entry.CreatedAt, check.CreatedAt)
}

func TestListEntry(t *testing.T) {
	for range 5 {
		createRandomEntry(t)
	}
	args := db.ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := TestQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, entries, 5)
}

func TestUpdateEntry(t *testing.T) {
	entry := createRandomEntry(t)

	arg := db.UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.CreateRandomMoney(),
	}

	update, err := TestQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)

	require.NotEqual(t, entry.Amount, update.Amount)
	require.Equal(t, entry.ID, update.ID)
	require.Equal(t, entry.AccountID, update.AccountID)
	require.Equal(t, entry.CreatedAt, update.CreatedAt)
	require.NotZero(t, update.ID)
	require.NotZero(t, update.CreatedAt)
}
