package lib

import (
	"upper.io/db.v3/mysql"
	"upper.io/db.v3"
	"github.com/labstack/gommon/log"
)

var configuracao = mysql.ConnectionURL{
	Host:"localhost",
	User: "root",
	Password:"root",
	Database:"membros",
}

var Conn db.Database
func init() {
	var err error

	Conn, err = mysql.Open(configuracao)
    if err != nil {
    	log.Fatal(err.Error())
	}
}
