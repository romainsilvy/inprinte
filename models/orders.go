package models

type OrderLine struct {
	name			string	`json:"name"`
	quantity		int		`json:"quantity"`
	photo			string	`json:"photo"`
}

type Order struct {
	id_order		int		   	`json:"id_order"`
	id_user			int		   	`json:"id_user"`
	date_order		string   	`json:"date_order"`
	order_state		string   	`json:"order_state"`
	price			int		   	`json:"price"`
	order_line		[]OrderLine	`json:"order_line"`

}

type AllOrders struct {
	allOrders	[]Order `json:"allOrders"`
}