package structures

type UserData struct {
	Email     string  `json:"Email"`
	Password  string  `json:"Password"`
	Firstname string  `json:"Firstname"`
	Lastname  string  `json:"Lastname"`
	Phone     string  `json:"Phone"`
	Address   Address `json:"Address"`
}

type UserProductData struct {
	Picture string `json:"Picture"`
	Name    string `json:"Name"`
	Price   string `json:"Price"`
}

type JsonResponseUsers struct {
	Type       string            `json:"type"`
	UserData   []UserData        `json:"UserData"`
	Favorite   []UserProductData `json:"Favorite"`
	OldCommand []UserProductData `json:"Oldcommand"`
	Message    string            `json:"message"`
}
