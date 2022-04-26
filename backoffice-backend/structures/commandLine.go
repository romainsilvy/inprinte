package structures

type GetCommandLines struct {
	Id         int    `json:"id"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Id_product int    `json:"id_product"`
	Id_command int    `json:"id_command"`
	Id_user    int    `json:"id_user"`
	Quantity   int    `json:"quantity"`
	Date       string `json:"date"`
	Price      int    `json:"price"`
	State      string `json:"status"`
	Name       string `json:"name"`
}

type GetCommandLine struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type UpdateCommandLine struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type CreateCommandLine struct {
	Id_product int    `json:"id_product,string"`
	Id_command int    `json:"id_command,string"`
	Quantity   int    `json:"quantity,string"`
	State      string `json:"state"`
}

type JsonReponseCommandLines struct {
	Id      int               `json:"id"`
	Type    string            `json:"type"`
	Data    CreateCommandLine `json:"data"`
	Message string            `json:"message"`
}

type JsonList struct {
	Id      int               `json:"id"`
	Type    string            `json:"type"`
	Data    []GetCommandLines `json:"data"`
	Message string            `json:"message"`
}
