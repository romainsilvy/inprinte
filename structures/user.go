package structures

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zipCode"`
}

type AllUsers struct {
	Id        int     `json:"id"`
	Email     string  `json:"email"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Phone     string  `json:"phone"`
	IsAlive   bool    `json:"is_alive"`
	Role      string  `json:"role"`
	Address   Address `json:"address"`
}

type OneUser struct {
	Id        int     `json:"id"`
	Email     string  `json:"email"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Phone     string  `json:"phone"`
	IsAlive   bool    `json:"is_alive"`
	Role      string  `json:"role"`
	Address   Address `json:"address"`
}
