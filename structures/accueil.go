package structures

type Accueil struct {
	Id_product  int    `json:"id"`
	Name        string `json:"Name"`
	Price       int    `json:"Price"`
	Description string `json:"Description"`
	Picture     string `json:"Picture"`
}

type JsonResponseAccueil struct {
	Type    string    `json:"type"`
	Data    []Accueil `json:"data"`
	Message string    `json:"message"`
}
