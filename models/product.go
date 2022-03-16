package models

type Product struct {
	id_product		int   `json:"id_product"`
	name			string   `json:"name"`
	description		string   `json:"description"`
	price			int   `json:"price "`
	category		string   `json:"category"`
	stars_number	int   `json:"stars_number"`
	photos			[]string `json:"photos"`

}

type AllProducts struct {
	allProducts		[]Product `json:"allProducts"`
}