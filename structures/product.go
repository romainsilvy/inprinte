package structures

type Product struct {
	Id_product   int     `json:"id_product"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        int     `json:"price "`
	Stars_number float64 `json:"stars_number"`
	Picture_url  string  `json:"picture_url"`
	Product_file string  `json:"product_file"`
}

type JsonResponseProduct struct {
	Type    string    `json:"type"`
	Data    []Product `json:"data"`
	Message string    `json:"message"`
}
