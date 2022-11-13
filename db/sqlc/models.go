// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
)

type Account struct {
	ID        int64          `json:"id"`
	Owner     string         `json:"owner"`
	Balance   sql.NullInt64  `json:"balance"`
	Currency  sql.NullString `json:"currency"`
	CreatedAt sql.NullTime   `json:"created_at"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be positive or negative
	Amount    sql.NullInt64 `json:"amount"`
	CreatedAt sql.NullTime  `json:"created_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be positive
	Amount    sql.NullInt64 `json:"amount"`
	CreatedAt sql.NullTime  `json:"created_at"`
}