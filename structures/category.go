package structures

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type OneCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type InsertOneResponse struct {
	Status string `json:"status"`
}
