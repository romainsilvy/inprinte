package structures

type Roles struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}

type OneRole struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}

type CreateOneRole struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}

type InsertOneRole struct {
	Id      int           `json:"id"`
	Type    string        `json:"type"`
	Data    CreateOneRole `json:"data"`
	Message string        `json:"message"`
}
