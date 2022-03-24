package CRUD

import (
	"database/sql"
	"inprinte/backend/utils"
)

func getCategories(db *sql.DB, sqlQuery string) []string {
	rows, err := db.Query(sqlQuery)

	// check errors
	utils.CheckErr(err)
	var categories []string

	// Foreach product
	for rows.Next() {
		var name string

		err = rows.Scan(&name)

		// check errors
		utils.CheckErr(err)

		categories = append(categories, name)
	}
	return categories
}
