package cookie

// import (
// 	"encoding/json"
// 	"inprinteBackoffice/structures"
// 	"inprinteBackoffice/utils"
// 	"net/http"
// )

// func Authenticate(w http.ResponseWriter, r *http.Request) {
// 	//set cors headers
// 	utils.SetCorsHeaders(&w)

// 	if r.Method == "POST" {
// 		var auth structures.GetAuthData
// 		var password string
// 		var is_alive bool
// 		var id_role int

// 		//decode json
// 		err := json.NewDecoder(r.Body).Decode(&auth)
// 		utils.CheckErr(err)

// 		//connect the database
// 		db := utils.DbConnect()

// 		sqlQuery := ("SELECT password, is_alive, id_role FROM user WHERE email = '" + auth.Email + "'")
// 		rows, err := db.Query(sqlQuery)
// 		utils.CheckErr(err)

// 		//parse the query
// 		for rows.Next() {
// 			//retrieve the values and check errors
// 			err = rows.Scan(&password, &is_alive, &id_role)
// 			utils.CheckErr(err)
// 		}

// 		// close the database connection
// 		db.Close()

// }
