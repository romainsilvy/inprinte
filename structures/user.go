package structures

type UserData struct {
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Phone     string  `json:"phone"`
	Address   Address `json:"address"`
}
type JsonResponseUsers struct {
	Type    string     `json:"type"`
	Data    []UserData `json:"data"`
	Message string     `json:"message"`
}
