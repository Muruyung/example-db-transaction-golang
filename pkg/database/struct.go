package database

import "database/sql"

// InitDB initialize database wrapper
type InitDB struct {
	ResellerDB *sql.DB
	TrxDB      *sql.DB
	ProductDB  *sql.DB
}
