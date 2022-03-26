package structures

type AllCommands struct {
	Id        int     `json:"id"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Id_user  int  `json:"id_user"`
	Price  int  `json:"price"`
	Quantity int  `json:"quantity"`
	Date string  `json:"date"`
	Status   string    `json:"status"`
}

type OneCommand struct {
	Id        int     `json:"id"`
	Status   string    `json:"status"`
}
