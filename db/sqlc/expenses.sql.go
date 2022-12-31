// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: expenses.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const deleteExpenses = `-- name: DeleteExpenses :exec
DELETE FROM expenses
WHERE id = $1
`

func (q *Queries) DeleteExpenses(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteExpenses, id)
	return err
}

const getExpenses = `-- name: GetExpenses :one
SELECT id, title, amount, note, tags FROM expenses
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetExpenses(ctx context.Context, id int32) (Expense, error) {
	row := q.db.QueryRowContext(ctx, getExpenses, id)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Amount,
		&i.Note,
		pq.Array(&i.Tags),
	)
	return i, err
}

const insertExpenses = `-- name: InsertExpenses :one
INSERT INTO expenses (
    title, amount, note, tags
) VALUES (
             $1, $2, $3, $4
         ) RETURNING id, title, amount, note, tags
`

type InsertExpensesParams struct {
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func (q *Queries) InsertExpenses(ctx context.Context, arg InsertExpensesParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, insertExpenses,
		arg.Title,
		arg.Amount,
		arg.Note,
		pq.Array(arg.Tags),
	)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Amount,
		&i.Note,
		pq.Array(&i.Tags),
	)
	return i, err
}

const listExpenses = `-- name: ListExpenses :many
SELECT id, title, amount, note, tags FROM expenses
ORDER BY id
`

func (q *Queries) ListExpenses(ctx context.Context) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, listExpenses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expense
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Amount,
			&i.Note,
			pq.Array(&i.Tags),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateExpenses = `-- name: UpdateExpenses :one
UPDATE expenses SET title = $2, amount = $3, note = $4, tags = $5 WHERE id = $1 RETURNING id, title, amount, note, tags
`

type UpdateExpensesParams struct {
	ID     int32    `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func (q *Queries) UpdateExpenses(ctx context.Context, arg UpdateExpensesParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, updateExpenses,
		arg.ID,
		arg.Title,
		arg.Amount,
		arg.Note,
		pq.Array(arg.Tags),
	)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Amount,
		&i.Note,
		pq.Array(&i.Tags),
	)
	return i, err
}
