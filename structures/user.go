package structures

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zipCode"`
}

type User struct {
	Id        int     `json:"id"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Phone     string  `json:"phone"`
	Address   Address `json:"address"`
}
