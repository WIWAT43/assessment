package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsertExpenses(t *testing.T) {
	arg := InsertExpensesParams{
		Title:  "Test",
		Amount: 21,
		Note:   "Test Note",
		Tags:   []string{"ABC", "DEF", "IJK"},
	}

	expense, err := testQueries.InsertExpenses(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, expense)

	require.Equal(t, arg.Title, expense.Title)
	require.Equal(t, arg.Amount, expense.Amount)
	require.Equal(t, arg.Note, expense.Note)
	require.Equal(t, arg.Tags, expense.Tags)

}
