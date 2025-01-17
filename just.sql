-- name: ListAccounts :many
SELECT id, owner, balance, currency, created_at
FROM accounts
ORDER BY id;

