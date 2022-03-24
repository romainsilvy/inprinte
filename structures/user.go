package structures

type UserData struct {
	Email     string  `json:"Email"`
	Password  string  `json:"Password"`
	Firstname string  `json:"Firstname"`
	Lastname  string  `json:"Lastname"`
	Phone     string  `json:"Phone"`
	Address   Address `json:"Address"`
}

type UserFavoriteProduct struct {
	Picture string `json:"Picture"`
	Name    string `json:"Name"`
	Price   string `json:"Price"`
}

type UserOldCommand struct {
	Picture  string `json:"Picture"`
	Name     string `json:"Name"`
	Price    string `json:"Price"`
	Quantity int    `json:"Quantity"`
}
type JsonResponseUsers struct {
	UserData        []UserData            `json:"UserData"`
	FavoriteProduct []UserFavoriteProduct `json:"FavoriteProduct"`
	OldCommand      []UserOldCommand      `json:"UserOldCommand"`
}
