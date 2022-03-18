package CRUD

import (
	_ "github.com/go-sql-driver/mysql"
)

/*func updateUser(id_user int, firstName, lastName, email, password, phone string, id_address, id_photo int, db *sql.DB) {
	_, err := db.Exec(`UPDATE users SET firstname = ?, lastname = ?, email = ?, password = ?, phone = ?, id_address = ?, id_photo = ? WHERE id = ?`, firstName, lastName, email, password, phone, id_address, id_photo, id_user)
	if err != nil {
		log.Fatal(err)
	}
}*/
