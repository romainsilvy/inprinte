package structures

type GetCategories struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	IsAlive bool   `json:"is_alive,string"`
}

type GetCategory struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	IsAlive bool   `json:"is_alive,string"`
}

type UpdateCategory struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	IsAlive bool   `json:"is_alive,string"`
}

type CreateCategory struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	IsAlive bool   `json:"is_alive,string"`
}

type JsonResponseCategory struct {
	Id      int            `json:"id"`
	Type    string         `json:"type"`
	Data    CreateCategory `json:"data"`
	Message string         `json:"message"`
}

type CategoriesList struct {
	CategoriesList []string `json:"CategoriesList"`
}
