package structures

type GetCommands struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Id_user   int    `json:"id_user"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	Date      string `json:"date"`
	Status    string `json:"status"`
}

type GetCommand struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type CreateCommand struct {
	Id_user int    `json:"id_user,string"`
	Date    string `json:"date"`
	State   string `json:"state"`
}

type UpdateCommand struct {
	Id_user int    `json:"id_user,string"`
	Date    string `json:"date"`
	State   string `json:"state"`
}
type JsonResponseCommand struct {
	Id      int           `json:"id"`
	Type    string        `json:"type"`
	Data    CreateCommand `json:"data"`
	Message string        `json:"message"`
}
