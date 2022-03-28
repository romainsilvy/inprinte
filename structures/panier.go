package structures

type CartItem struct {
	Id_product  int     `json:"id_product"`
	Picture     string  `json:"picture"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type Cart struct {
	AllItems []CartItem `json:"allItems"`
}
