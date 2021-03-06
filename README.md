# go-logstock

Tool to check duplicate query to PostgreSQL database based on [go-pg ORM](https://github.com/go-pg/pg).


[![Build Status](https://travis-ci.org/revdaalex/go-logstock.svg?branch=master)](https://travis-ci.org/revdaalex/go-logstock)
[![Go Report Card](https://goreportcard.com/badge/github.com/revdaalex/go-logstock)](https://goreportcard.com/report/github.com/revdaalex/go-logstock)

Import:

```go
. "github.com/revdaalex/go-logstock/src/go_logstock"
```
go get github.com/revdaalex/go-logstock

Use in test:

```go
func TestQuery(t *testing.T)  {
	db, err := DBConn(pgOptions)
	if err != nil {
		panic(err)
	}
	// check func with query
	testQuery(db)
	// Create log and assert log query and db query
	CheckLog(t, "logName")
}
```

Log files are created in the log directory GOPATH/log