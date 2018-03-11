package go_logstock

import (
	"testing"
	_"fmt"
	"github.com/go-pg/pg"
)

type User struct {
	ID   int64
	Name string
}

func TestConnect(t *testing.T) {
	db, err := DBConn(pgOptions())
	if err != nil {
		panic(err)
	}
	Create(db)
	defer db.Close()
	//fmt.Println(TestQuery)
	CheckLog(t, "test.log")
}

func Create(db *pg.DB) {
	var test string
	db.Exec(userTableSQL)
	_, err := db.Query(pg.Scan(&test), "SELECT value FROM public.user WHERE id=?", 1)
	if err != nil {
		panic(err)
	}
}
