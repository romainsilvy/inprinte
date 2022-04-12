package structures

type MostSales struct {
	Id_product int     `json:"id"`
	Name       string  `json:"Name"`
	Price      float64 `json:"Price"`
	Picture    string  `json:"Picture"`
}

type BestRatedProduct struct {
	Id_product int     `json:"id"`
	Name       string  `json:"Name"`
	Price      float64 `json:"Price"`
	Picture    string  `json:"Picture"`
}

type Accueil struct {
	MostSales []MostSales        `json:"mostSales"`
	BestRated []BestRatedProduct `json:"bestRated"`
}
