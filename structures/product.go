package structures

type ProductData struct {
	Id_product   int     `json:"id_product"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price "`
	Picture      string  `json:"picture"`
	Product_file string  `json:"product_file"`
}

type Product struct {
	ProductData ProductData `json:"product_data"`
}
