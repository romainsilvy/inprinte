package structures

type AllRates struct {
	Id           int    `json:"id"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Id_product   int    `json:"id_product"`
	Stars_number int    `json:"stars_number,string"`
	Id_user      int    `json:"id_user"`
	Name         string `json:"product_name"`
}

type OneRate struct {
	Id           int `json:"id"`
	Stars_number int `json:"stars_number,string"`
}

type CreateOneRate struct {
	Id_user      int `json:"id_user,string"`
	Id_product   int `json:"id_product,string"`
	Stars_number int `json:"stars_number,string"`
}

type InsertOneRate struct {
	Id      int           `json:"id"`
	Type    string        `json:"type"`
	Data    CreateOneRate `json:"data"`
	Message string        `json:"message"`
}
