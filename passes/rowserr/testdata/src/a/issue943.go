package a

import (
	"database/sql"
	"fmt"
	"math/rand"
)

var X int

func issue943_1() {
	db, err := sql.Open("postgres", "postgres://localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var rows *sql.Rows
	if rand.Float64() < 0.5 {
		rows, err = db.Query("select 1")
	} else {
		rows, err = db.Query("select 2")
	}
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		fmt.Println("new rows")
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
}

func issue943_2() {
	db, err := sql.Open("postgres", "postgres://localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("select 1")
	defer rows.Close()
	if err := rows.Err(); err != nil {
		panic(err)
	}
}
