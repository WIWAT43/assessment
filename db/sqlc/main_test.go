package db

import (
	"assessment/config"
	"database/sql"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config.InitViper("../../")
	cf := config.GetConfig()
	conn, err := sql.Open(cf.DbDriver, cf.DbSource)
	if err != nil {
		log.Fatal("Can not connect to db: ", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
