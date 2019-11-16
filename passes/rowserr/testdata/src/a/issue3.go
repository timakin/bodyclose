package a

import (
	"database/sql"
	"fmt"
)

func issue3_2(db *sql.DB) {
	rows, _ := db.Query("")
	defer func() {
		_ = rows.Err()
	}()
}

func issue3_3(db *sql.DB) {
	rows, _ := db.Query("")
	defer func() { fmt.Println(rows.Err()) }()
}

func funcReceiver(msg string, er error) {
	fmt.Println(msg)
	if er != nil {
		fmt.Println(er)
	}
}

func issue3_4(db *sql.DB) {
	rows, _ := db.Query("")
	defer func() { funcReceiver("test", rows.Err()) }()
}
