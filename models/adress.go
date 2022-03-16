package models

type Adress struct {
	id_adress	int		`json:"id_adress"`
	street		string	`json:"street"`
	city		string	`json:"city"`
	state       string	`json:"state"`
	country     string	`json:"country "`
	zipCode     string	`json:"zipCode"`
}