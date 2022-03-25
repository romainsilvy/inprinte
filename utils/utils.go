package utils

import (
	"database/sql"
	"net/http"
	"strconv"
)

func SetCorsHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Max-Age", "86400")
}

func SetXTotalCountHeader(w *http.ResponseWriter, totalCount int) {
	(*w).Header().Set("Access-Control-Expose-Headers", "X-Total-Count")
	(*w).Header().Set("X-Total-Count", strconv.Itoa(totalCount))
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func DbConnect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "inprinteAdmin"
	dbPass := "louisensueur"
	dbProtocol := "tcp"
	dbIp := "178.170.14.134"
	dbPort := "3306"
	dbName := "inprinte"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbProtocol+"("+dbIp+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
