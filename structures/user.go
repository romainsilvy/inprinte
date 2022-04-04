package structures

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zipCode"`
}

type GetUsers struct {
	Id        int     `json:"id"`
	Email     string  `json:"email"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Phone     string  `json:"phone"`
	IsAlive   bool    `json:"is_alive"`
	Role      string  `json:"role"`
	Address   Address `json:"address"`
}

type GetUser struct {
	Id        int     `json:"id"`
	Email     string  `json:"email"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Phone     string  `json:"phone"`
	IsAlive   bool    `json:"is_alive"`
	Role      string  `json:"role"`
	Address   Address `json:"address"`
}

type CreateUser struct {
	Email     string  `json:"email"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Password  string  `json:"password"`
	Phone     string  `json:"phone"`
	IsAlive   bool    `json:"is_alive,string"`
	Role      string  `json:"role"`
	Address   Address `json:"address"`
}

type UpdateUser struct {
	Email     string  `json:"email"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Password  string  `json:"password"`
	Phone     string  `json:"phone"`
	IsAlive   bool    `json:"is_alive,string"`
	Role      string  `json:"role"`
	Address   Address `json:"address"`
}

type JsonResponseUser struct {
	Id      int        `json:"id"`
	Type    string     `json:"type"`
	Data    CreateUser `json:"data"`
	Message string     `json:"message"`
}
