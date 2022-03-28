package structures

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zipCode"`
}

type UserData struct {
	Id        int     `json:"id"`
	Email     string  `json:"email"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Phone     string  `json:"phone"`
	Address   Address `json:"address"`
}

type UserCommandHistory struct {
	Id       int     `json:"id"`
	Picture  string  `json:"picture"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type UserFavoriteProducts struct {
	Id      int     `json:"idProduct"`
	Picture string  `json:"picture"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
}

type User struct {
	UserData         UserData               `json:"userData"`
	FavoriteProducts []UserFavoriteProducts `json:"favoriteProducts"`
	CommandHistory   []UserCommandHistory   `json:"userCommandHistory"`
}
