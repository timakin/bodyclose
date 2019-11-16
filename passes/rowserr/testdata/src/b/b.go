package b

import (
	"database/sql"
	"io"
)

func RowsErrCheck(db *sql.DB) {
	rows, _ := db.Query("")

	var i io.ReadCloser
	i.Close()

	defer func() {
		_ = rows.Err()
	}()

}

func get(db *sql.DB) *sql.Rows {
	rows, _ := db.Query("")
	return rows
}

func xx() {
	resp := get(new(sql.DB))
	_ = resp.Err()
}

func xxx(db *sql.DB) {
	rows, _ := db.Query("")
	_ = rows.Err()
}
