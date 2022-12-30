-- name: GetExpenses :one
SELECT * FROM expenses
WHERE id = ? LIMIT 1;

-- name: ListExpenses :many
SELECT * FROM expenses
ORDER BY id;

-- name: InsertExpenses :one
INSERT INTO expenses (
    title, amount, note, tags
) VALUES (
             ?, ?, ?, ?
         ) RETURNING *;

-- name: UpdateExpenses :one
UPDATE expenses SET title = ?, amount = ?, note = ?, tags = ? WHERE id = ? RETURNING *;

-- name: DeleteExpenses :exec
DELETE FROM expenses
WHERE id = ?;