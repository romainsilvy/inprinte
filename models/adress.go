package models

type Adress struct {
	street		 string   `json:"street"`
	city         string   `json:"city"`
	state        string   `json:"state"`
	country      string   `json:"country "`
	zipCode      string   `json:"zipCode"`
}