package b

import (
	"database/sql"
)

func f10(db *sql.DB) {

	rows, _ := db.Query("")
	 func(rs *sql.Rows) {
		_ = rs.Err()
	}(rows)
	//resCloser(rows)
}
