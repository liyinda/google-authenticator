package sqlite

import (
	"database/sql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var (
	authEngine *xorm.Engine
)

func init() {
	if f, _ := fileExists("./authenticator.db"); f == false {
		createTable()
	}
	//init google-authenticator db
	var err error
	authEngine, err = xorm.NewEngine("sqlite3", "./authenticator.db")
	authEngine.ShowSQL(true)
	if err != nil {
		log.Fatal(err)
	}
	falarmcenter, err := os.Create("authenticator.log")
	if err != nil {
		println(err.Error())
		return
	}
	authEngine.SetLogger(xorm.NewSimpleLogger(falarmcenter))
	err = authEngine.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connect postgresql google-authenticator success")
}

func createTable() {
	//open database, if table is't not exist, create it
	db, err := sql.Open("sqlite3", "./authenticator.db")
	log.Printf("store.sqlite.init:error:%v", err)

	sql_table := `
	CREATE TABLE IF NOT EXISTS authuser(
	id integer PRIMARY KEY autoincrement,
	name varchar NOT NULL,
	dr int4 NOT NULL DEFAULT 0,
	secretid varchar NULL,
	qrcode varchar NOT NULL,
	createdtime varchar NULL
	);
	`
	db.Exec(sql_table)
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
