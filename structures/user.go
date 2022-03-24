package structures

type User struct {
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Phone     string  `json:"phone"`
	Address   Address `json:"address"`
}
type JsonResponseUsers struct {
	Data []User `json:"data"`
}
