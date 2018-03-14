package go_logstock

import (
	"github.com/go-pg/pg"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
)

var (
	testArray []string
	testQuery string
	logQuery  string
	logArray  []string
)

func DBConn(opt *pg.Options) (*pg.DB, error) {

	db := pg.Connect(opt)

	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}
		testArray = append(testArray, query)
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
	testQuery = strings.Join(testArray, "")

	logArray = strings.Split(logQuery, " ")
	testArray = strings.Split(testQuery, " ")

	if !assert.Equal(t, logArray, testArray) {
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

	for _, f := range testArray {
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
