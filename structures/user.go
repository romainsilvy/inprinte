package structures

type UserData struct {
	Id        int     `json:"id"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Phone     string  `json:"phone"`
	Address   Address `json:"address"`
}

type UserFavoriteProduct struct {
	Picture string `json:"picture"`
	Name    string `json:"name"`
	Price   string `json:"price"`
}

type UserOldCommand struct {
	Picture  string `json:"picture"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity int    `json:"quantity"`
}
type JsonResponseUsers struct {
	Data            []UserData            `json:"data"`
	FavoriteProduct []UserFavoriteProduct `json:"favoriteProduct"`
	OldCommand      []UserOldCommand      `json:"userOldCommand"`
}
