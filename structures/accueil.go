package structures

type Accueil struct {
	Id_product  int    `json:"id"`
	Name        string `json:"Name"`
	Price       int    `json:"Price"`
	Description string `json:"Description"`
	Picture     string `json:"Picture"`
}

type JsonResponseAccueil struct {
	Data []Accueil `json:"data"`
}
