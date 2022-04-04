package structures

type GetRoles struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}

type GetRole struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}

type CreateRole struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}

type UpdateRole struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}

type JsonResponseRole struct {
	Id      int        `json:"id"`
	Type    string     `json:"type"`
	Data    CreateRole `json:"data"`
	Message string     `json:"message"`
}
