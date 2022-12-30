package db

import (
	"assessment/util"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func createExpenses(t *testing.T) Expense {
	arg := InsertExpensesParams{
		Title:  util.RandomString(20),
		Amount: float64(util.RandomInt(10, 200)),
		Note:   util.RandomString(100),
		Tags:   []string{"ABC", "DEF", "IJK"},
	}

	expense, err := testQueries.InsertExpenses(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, expense)
	require.Equal(t, arg.Title, expense.Title)
	require.Equal(t, arg.Amount, expense.Amount)
	require.Equal(t, arg.Note, expense.Note)
	require.Equal(t, arg.Tags, expense.Tags)

	return expense
}

func TestQueries_InsertExpenses(t *testing.T) {
	createExpenses(t)
}

func TestQueries_UpdateExpenses(t *testing.T) {

	exp1 := createExpenses(t)

	arg := UpdateExpensesParams{
		ID:     exp1.ID,
		Title:  util.RandomString(20),
		Amount: float64(util.RandomInt(10, 200)),
		Note:   util.RandomString(100),
		Tags:   []string{"ABC", "DEF", "IJK"},
	}

	expense, err := testQueries.UpdateExpenses(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, expense)
	require.Equal(t, arg.Title, expense.Title)
	require.Equal(t, arg.Amount, expense.Amount)
	require.Equal(t, arg.Note, expense.Note)
	require.Equal(t, arg.Tags, expense.Tags)

}
func TestQueries_GetExpenses(t *testing.T) {
	exp1 := createExpenses(t)
	exp2, err := testQueries.GetExpenses(context.Background(), exp1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, exp2)
	require.Equal(t, exp1.Title, exp2.Title)
	require.Equal(t, exp1.Amount, exp2.Amount)
	require.Equal(t, exp1.Note, exp2.Note)
	require.Equal(t, exp1.Tags, exp2.Tags)
}

func TestQueries_ListExpenses(t *testing.T) {
	arg := ListExpensesParams{
		Limit:  5,
		Offset: 5,
	}

	for i := 0; i < 10; i++ {
		createExpenses(t)
	}

	exp, err := testQueries.ListExpenses(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, exp, 5)

	for _, ex := range exp {
		require.NotEmpty(t, ex)
	}
}

func TestQueries_DeleteExpenses(t *testing.T) {
	exp := createExpenses(t)
	err := testQueries.DeleteExpenses(context.Background(), exp.ID)

	require.NoError(t, err)

	expCheck, err := testQueries.GetExpenses(context.Background(), exp.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, expCheck)
}
