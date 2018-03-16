package go_logstock

import (
	"testing"
	"github.com/go-pg/pg"
)

const userTableSQL = `
CREATE TABLE public.user (
id int, 
value text
);
INSERT INTO public.user VALUES (1, 'test')
`

func TestConnect(t *testing.T) {
	db, err := DBConn(pgOptions())
	if err != nil {
		panic(err)
	}
	Create(db)
	defer db.Close()
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

func pgOptions() *pg.Options {
	return &pg.Options{
		User:     "postgres",
		Database: "test",
		Password: "8777738",
	}
}
