package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/fatih/color"
)

func SetCorsHeaders(w *http.ResponseWriter) {

	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Max-Age", "86400")
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func InprinteAscii() {
	fmt.Println("  ___ _  _ ___ ___ ___ _  _ _____ ___")
	fmt.Println(" |_ _| \\| | _ \\ _ \\_ _| \\| |_   _| __|")
	fmt.Println("  | || .` |  _/   /| || .` | | | | _|")
	fmt.Println(" |___|_|\\_|_| |_|_\\___|_|\\_| |_| |___|\n\n	")
}

func PrettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}

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

func GetAllParams(r *http.Request) (string, string) {
	containsOrder := true
	containsSort := true
	containsStart := true
	containsEnd := true
	rangeBy := ""
	orderBy := ""

	urlOrder, ok := r.URL.Query()["_order"]
	if !ok || len(urlOrder[0]) < 1 {
		log.Println("Url Param 'order' is missing")
		containsOrder = false
	}

	urlSort, ok := r.URL.Query()["_sort"]
	if !ok || len(urlSort[0]) < 1 {
		log.Println("Url Param 'sort' is missing")
		containsSort = true
	}

	urlStart, ok := r.URL.Query()["_start"]
	if !ok || len(urlStart[0]) < 1 {
		log.Println("Url Param 'start' is missing")
		containsStart = false
	}

	urlEnd, ok := r.URL.Query()["_end"]
	if !ok || len(urlEnd[0]) < 1 {
		log.Println("Url Param 'End' is missing")
		containsEnd = false
	}

	if containsOrder && containsSort {
		orderBy = " ORDER BY " + urlSort[0] + " " + urlOrder[0]
	}

	if containsStart && containsEnd {
		rangeBy = " LIMIT " + urlStart[0] + "," + urlEnd[0]
	}
	return orderBy, rangeBy
}

func GetRangeParam(r *http.Request) string {
	containsStart := true
	containsEnd := true
	rangeBy := ""

	urlStart, ok := r.URL.Query()["_start"]
	if !ok || len(urlStart[0]) < 1 {
		log.Println("Url Param 'start' is missing")
		containsStart = false
	}

	urlEnd, ok := r.URL.Query()["_end"]
	if !ok || len(urlEnd[0]) < 1 {
		log.Println("Url Param 'End' is missing")
		containsEnd = false
	}

	if containsStart && containsEnd {
		rangeBy = " LIMIT " + urlStart[0] + "," + urlEnd[0]
	}
	return rangeBy
}
