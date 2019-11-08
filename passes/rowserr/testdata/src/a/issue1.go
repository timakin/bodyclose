package a

import "database/sql"

func get(db *sql.DB) *sql.Rows {
	rows, _ := db.Query("")
	return rows
}

func main() {
	resp := get(new(sql.DB))
	_ = resp.Err()
}
