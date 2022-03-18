package structures

type Accueil struct {
	Name   			string    `json:"Name"`
	Price     		string `json:"Price"`
	Description     string `json:"Description"`
	//Picture    	string `json:"Picture"`

}

type JsonResponseAccueil struct {
    Type    string 		`json:"type"`
    Data    []Accueil 	`json:"data"`
    Message string 		`json:"message"`
}