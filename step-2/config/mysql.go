package config

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

const (
	User     string = "root"
	Password string = ""
	Db       string = "fantasy"
	Host     string = "localhost"
	Port     string = "3306"
)

func GetSession() *dbr.Session {

	db, err := dbr.Open("mysql", User+":"+Password+"@tcp("+Host+":"+Port+")/"+Db+"?parseTime=true&loc=Asia%2FTokyo", nil)

	if err != nil {
		logrus.Error(err)
	} else {
		session := db.NewSession(nil)
		return session
	}
	return nil
}