package structures

type Product struct {
	Id_product   int      `json:"id_product"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Price        int      `json:"price "`
	Category     string   `json:"category"`
	Stars_number int      `json:"stars_number"`
	Photos       []string `json:"photos"`
}

type AllProducts struct {
	AllProducts []Product `json:"allProducts"`
}
