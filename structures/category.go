package structures

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type OneCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateOneCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type InsertOneCategory struct {
	Id      int               `json:"id"`
	Type    string            `json:"type"`
	Data    CreateOneCategory `json:"data"`
	Message string            `json:"message"`
}
