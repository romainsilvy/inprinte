package structures

type BoutiqueProduct struct {
	Id_product  int    `json:"Id_product"`
	Name        string `json:"Name"`
	Price       string `json:"Price"`
	Description string `json:"Description"`
	Picture     string `json:"Picture"`
}

type JsonResponseBoutique struct {
	BoutiqueNews       []BoutiqueProduct `json:"BoutiqueNews"`
	BoutiqueMostWanted []BoutiqueProduct `json:"BoutiqueMostWanted"`
	AllProducts        []BoutiqueProduct `json:"AllProducts"`
	Categories         []string          `json:"Categories"`
}
