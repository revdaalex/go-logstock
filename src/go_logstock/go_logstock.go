package go_logstock

import (
	"github.com/go-pg/pg"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
)

var (
	testQuery string
	testArray []string
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

func CheckLog(t TestingT, logName string) {
	createDir()
	readLog(logName)
	testQuery = strings.Join(testArray, "")
	testArray = strings.Split(testQuery, "\n")

	if !assert.Equal(t, logArray, testArray) {
		t.FailNow()
	}
}

func createDir() {
	err := os.Mkdir(os.Getenv("GOPATH")+"/log", 0755)
	if err != nil {
		return
	}
}

func createLog(logName string) {
	file, err := os.Create(os.Getenv("GOPATH") + "/log/" + logName + ".log")
	if err != nil {
		return
	}
	defer file.Close()

	for _, f := range testArray {
		file.WriteString(f)
	}
}

func readLog(logName string) {
	bs, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/log/" + logName + ".log")
	if err != nil && os.IsNotExist(err) {
		createLog(logName)
		readLog(logName)
	}
	if len(logArray) == 0 {
		logArray = strings.Split(string(bs), "\n")
	}
}
