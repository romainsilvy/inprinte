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
	State      string `json:"state"`
	Name       string `json:"name"`
}
