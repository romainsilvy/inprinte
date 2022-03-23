package databaseTools

import (
	"database/sql"

	"github.com/fatih/color"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "inprinteAdmin"
	dbPass := "louisensueur"
	dbProtocol := "tcp"
	dbIp := "178.170.14.134"
	dbPort := "3306"
	dbName := "inprinte"
	color.Green("Connecting database ...")

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbProtocol+"("+dbIp+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	} else {
		color.Cyan("Database connected !\n")
	}
	return db
}
