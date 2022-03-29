package structures

type FileUrl struct {
	Url string `json:"url"`
}

type PictureUrl struct {
	Url string `json:"url"`
}

type AllProducts struct {
	Id                 int          `json:"id"`
	Name               string       `json:"name"`
	Price              int          `json:"price"`
	Description        string       `json:"description"`
	Pending_validation bool         `json:"pending_validation"`
	Is_alive           bool         `json:"is_alive"`
	Category           string       `json:"category"`
	Firstname          string       `json:"firstname"`
	Lastname           string       `json:"lastname"`
	Role               string       `json:"role"`
	Id_user            int          `json:"id_user"`
	Rate               float64      `json:"rate"`
	FileUrl            []FileUrl    `json:"fileUrl"`
	PictureUrl         []PictureUrl `json:"pictureUrl"`
}

type OneProduct struct {
	Id                 int          `json:"id"`
	Name               string       `json:"name"`
	Price              int          `json:"price"`
	Description        string       `json:"description"`
	Pending_validation bool         `json:"pending_validation"`
	Is_alive           bool         `json:"is_alive"`
	Category           string       `json:"category"`
	Firstname          string       `json:"firstname"`
	Lastname           string       `json:"lastname"`
	Role               string       `json:"role"`
	Id_user            int          `json:"id_user"`
	Rate               float64      `json:"rate"`
	FileUrl            []FileUrl    `json:"fileUrl"`
	PictureUrl         []PictureUrl `json:"pictureUrl"`
}
