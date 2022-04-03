package structures

type AllCommandLines struct {
	Id         int    `json:"id"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Id_product int    `json:"id_product"`
	Id_command int    `json:"id_command"`
	Id_user    int    `json:"id_user"`
	Quantity   int    `json:"quantity"`
	Date       int    `json:"date"`
	Price      int    `json:"price"`
	State      string `json:"status"`
	Name       string `json:"name"`
}

type OneCommandLines struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type JsonReponseCommandLines struct {
	Data   []OneCommandLines `json:"data"`
	Id     int               `json:"id"`
	Status string            `json:"status"`
}

type CreateOneCommandLines struct {
	Id_product int    `json:"id_product,string"`
	Id_command int    `json:"id_command,string"`
	Quantity   int    `json:"quantity,string"`
	State      string `json:"state"`
}

type InsertOneCommandLines struct {
	Id      int                   `json:"id"`
	Type    string                `json:"type"`
	Data    CreateOneCommandLines `json:"data"`
	Message string                `json:"message"`
}
