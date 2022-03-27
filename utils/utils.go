package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"

	"github.com/fatih/color"
)

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
