package models

type Adress struct {
	Id_adress int    `json:"id_adress"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country "`
	ZipCode   string `json:"zipCode"`
}
