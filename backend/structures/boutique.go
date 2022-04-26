package structures

type NewProduct struct {
	Id_product int     `json:"Id_product"`
	Name       string  `json:"Name"`
	Price      float64 `json:"Price"`
	Picture    string  `json:"Picture"`
}
type BoutiqueProduct struct {
	Id_product  int     `json:"Id_product"`
	Name        string  `json:"Name"`
	Price       float64 `json:"Price"`
	Picture     string  `json:"Picture"`
	Description string  `json:"Description"`
}

type Boutique struct {
	NewProducts []NewProduct      `json:"BoutiqueNews"`
	MostSales   []MostSales       `json:"BoutiqueMostWanted"`
	AllProducts []BoutiqueProduct `json:"AllProducts"`
	Categories  []string          `json:"Categories"`
}
