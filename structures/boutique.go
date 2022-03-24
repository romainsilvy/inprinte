package structures

type BoutiqueProduct struct {
	Name        string `json:"Name"`
	Price       string `json:"Price"`
	Description string `json:"Description"`
	Picture     string `json:"Picture"`
}

type BoutiqueCategories struct {
	GlobalCategory string   `json:"GlobalCategory"`
	SubCategory    []string `json:"SubCategory"`
}

type JsonResponseBoutique struct {
	Type               string               `json:"type"`
	BoutiqueNews       []BoutiqueProduct    `json:"BoutiqueNews"`
	BoutiqueMostWanted []BoutiqueProduct    `json:"BoutiqueMostWanted"`
	AllProducts        []BoutiqueProduct    `json:"AllProducts"`
	Categories         []BoutiqueCategories `json:"Categories"`
	Message            string               `json:"message"`
}
