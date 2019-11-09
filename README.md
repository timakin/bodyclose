# rowserrcheck

[![CircleCI](https://circleci.com/gh/jingyugao/rowserrcheck.svg?style=svg)](https://circleci.com/gh/jingyugao/rowserrcheck)

`rowserrcheck` is a static analysis tool which checks whether `sql.Rows.Err` is correctly checked.

## Install

You can get `rowserrcheck` by `go get` command.

```bash
$ go get -u github.com/jingyugao/rowserrcheck
```

## Analyzer

`rowserrcheck` validates whether [*database/sql.Rows](https://golang.org/pkg/database/sql/#Rows.Err) of sql query calls method `rows.Err()` such as below code.

```go
rows, _ := db.Query("select id from tb") // Wrong case
if err != nil {
	// handle error
}
for rows.Next(){
	// handle rows
}
```

This code is wrong. You must check rows.Err when finished scan rows.

```go
rows, _ := db.Query("select id from tb") // Wrong case
for rows.Next(){
	// handle rows
}
if rows.Err()!=nil{
	// handle err
}
```

In the [GoDoc of sql.Rows](https://golang.org/pkg/database/sql/#Rows) this rule is clearly described.

If you forget this sentence, and unluckly an `invaliad connection` error happend when fetch
data from database, `rows.Next` will return false, and you will get an incomplete data, and
even it seems everything is ok. This will cause serious accident.

## Thanks
Thanks for [timakin](https://github.com/jingyugao/rowserrcheck).