package postgres

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/liyinda/google-authenticator/utils/apple"
	"log"
	"os"
)

var (
	authEngine *xorm.Engine
)

const (
	host            = ""
	port            = 5432
	user            = "postgres"
	authenticatorDB = "authenticator"
)

func init() {
	//get apple decode
	key := apple.GetAppleKey()
	decode := apple.DeMD5Code("", key)
	//init google-authenticator db
	var err error
	authSqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, decode, authenticatorDB)
	authEngine, err = xorm.NewEngine("postgres", authSqlInfo)
	authEngine.ShowSQL(true)
	if err != nil {
		log.Fatal(err)
	}
	falarmcenter, err := os.Create("alarmcenterengine.log")
	if err != nil {
		println(err.Error())
		return
	}
	authEngine.SetLogger(xorm.NewSimpleLogger(falarmcenter))
	err = authEngine.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connect postgresql alarm-center success")
}
