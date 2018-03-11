package go_logstock

import (
	"github.com/go-pg/pg"
	"io/ioutil"
	"os"
	"strings"
	"github.com/stretchr/testify/assert"
)

const userTableSQL = `
CREATE TABLE public.user (
id int, 
value text
);
INSERT INTO public.user VALUES (1, 'test')
`

var (
	arrayQuery []string
	testQuery  string
	logQuery   string
)

func DBConn(opt *pg.Options) (*pg.DB, error) {

	db := pg.Connect(opt)

	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}
		arrayQuery = append(arrayQuery, query)
	})
	return db, nil
}

func pgOptions() *pg.Options {
	return &pg.Options{
		User:     "postgres",
		Database: "test",
		Password: "8777738",
	}
}

func CheckLog(t TestingT, logName string) {
	createDir()
	readLog(logName)
	testQuery = strings.Join(arrayQuery, "")

	if !assert.Equal(t, logQuery, testQuery) {
		t.FailNow()
	}
}

func createDir() {
	err := os.Mkdir(os.Getenv("GOPATH")+"/log", os.ModePerm)
	if err != nil {
		return
	}
}

func createLog(logName string) {
	file, err := os.Create(os.Getenv("GOPATH") + "/log/" + logName)
	if err != nil {
		return
	}
	defer file.Close()

	for _, f := range arrayQuery {
		file.WriteString(f)
	}
}

func readLog(logName string) {
	bs, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/log/" + logName)
	if err != nil {
		createLog(logName)
		readLog(logName)
	}
	logQuery = string(bs)
}
