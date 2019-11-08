package b

import (
	"database/sql"
)

func f10(db *sql.DB) {

	rows, _ := db.Query("")
	resCloser := func(rs *sql.Rows) {
		_ = rs.Err()
	}
	resCloser(rows)
}
