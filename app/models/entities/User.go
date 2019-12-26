package entities

import (
	"database/sql"
)

type User struct {
	ID        int64
	UserName  string
	Password  string
	DBConnect *sql.DB
}
