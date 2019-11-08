package a

import (
	"database/sql"
)

func testNoCrashOnDefer(db *sql.DB) {
	rows, _ := db.Query("")
	for rows.Next() {
	}

	defer func(rs *sql.Rows) {
		_ = rs.Err()
	}(rows)
}
