package models

type Product struct {
	id_product		string   `json:"id_product"`
	name			string   `json:"name"`
	description		string   `json:"description"`
	price			string   `json:"price "`
	category		string   `json:"category"`
	stars_number	string   `json:"stars_number"`
	photos			[]string `json:"photos"`

}

type AllProducts struct {
	allProducts		[]User `json:"allProducts"`
}