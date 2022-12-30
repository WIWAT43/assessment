-- name: GetExpenses :one
SELECT * FROM expenses
WHERE id = $1 LIMIT 1;

-- name: ListExpenses :many
SELECT * FROM expenses
ORDER BY id LIMIT $1
OFFSET $2;

-- name: InsertExpenses :one
INSERT INTO expenses (
    title, amount, note, tags
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

-- name: UpdateExpenses :one
UPDATE expenses SET title = $2, amount = $3, note = $4, tags = $5 WHERE id = $1 RETURNING *;

-- name: DeleteExpenses :exec
DELETE FROM expenses
WHERE id = $1;