package utils

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}