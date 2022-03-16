package CRUD 

// import (
// 	"log"
// 	"database/sql"

// 	databaseTools "inprinte/backend/database"
// 	_ "github.com/go-sql-driver/mysql"
// )

// func Create() {
// 	db := databaseTools.DbConnect();
// 	table := "users"

// 	switch table {
// 	case "users":
// 		createUser("ismael", "Hacquin", "ismael.hacquin@gmail.com", "ismathegoat", "0909090909", 1, 1, db)
// 	case "products":
// 		createProduct("ismael", "Hacquin", "ismael.hacquin@gmail.com", "ismathegoat", "0909090909", 1, 1, db)
// 	}
// }

// func createUser(firstName, lastName, email, password, phone string, id_address, id_photo int, db *sql.DB) {
// 	_, err := db.Exec(`INSERT INTO users (firstname, lastname, email, password, phone, id_address, id_photo) VALUES (?, ?, ?, ?, ?, ?, ?)`, firstName, lastName, email, password, phone, id_address, id_photo)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }