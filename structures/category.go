package structures

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type OneCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type InsertOne struct {
	Id      int         `json:"id"`
	Type    string      `json:"type"`
	Data    OneCategory `json:"data"`
	Message string      `json:"message"`
}
