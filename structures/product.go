package structures

type FileUrl struct {
	Url string `json:"url"`
}

type PictureUrl struct {
	Url string `json:"url"`
}

type GetProducts struct {
	Id                 int     `json:"id"`
	Name               string  `json:"name"`
	Price              int     `json:"price"`
	Description        string  `json:"description"`
	Pending_validation bool    `json:"pending_validation"`
	Is_alive           bool    `json:"is_alive"`
	Category           string  `json:"category"`
	Firstname          string  `json:"firstname"`
	Lastname           string  `json:"lastname"`
	Role               string  `json:"role"`
	Id_user            int     `json:"id_user"`
	Rate               float64 `json:"rate"`
}

type GetProduct struct {
	Id                 int      `json:"id"`
	Name               string   `json:"name"`
	Price              int      `json:"price"`
	Description        string   `json:"description"`
	Pending_validation bool     `json:"pending_validation"`
	Is_alive           bool     `json:"is_alive"`
	Category           string   `json:"category"`
	Firstname          string   `json:"firstname"`
	Lastname           string   `json:"lastname"`
	Role               string   `json:"role"`
	Id_user            int      `json:"id_user"`
	Rate               float64  `json:"rate"`
	FileUrl            []string `json:"fileUrl"`
	PictureUrl         []string `json:"pictureUrl"`
}

type CreateProduct struct {
	Name               string `json:"name"`
	Price              int    `json:"price,string"`
	Description        string `json:"description"`
	Pending_validation bool   `json:"pending_validation,string"`
	Is_alive           bool   `json:"is_alive,string"`
	Category           string `json:"category"`
	Id_user            int    `json:"id_user,string"`
}

type JsonResponseProduct struct {
	Id      int           `json:"id"`
	Type    string        `json:"type"`
	Data    CreateProduct `json:"data"`
	Message string        `json:"message"`
}
