package models

type User struct {
	id_user		string		`json:"id_user"`
	photo		string		`json:"photo"`
	email		string		`json:"email"`
	password	string		`json:"password "`
	firstname	string		`json:"firstname"`
	lastname	string		`json:"lastname"`
	phone		string		`json:"phone"`
	adress		[]Adress 	`json:"adress"`
}

type AllUsers struct {
	allUsers	[]User		`json:"allUsers"`
}