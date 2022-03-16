package models

type OrderLine struct {
	name			string   `json:"name"`
	quantity		string   `json:"quantity"`
	photo			string   `json:"photo"`
}

type Order struct {
	id_order		string   	`json:"id_order"`
	id_user			string   	`json:"id_user"`
	date_order		string   	`json:"date_order"`
	order_state		string   	`json:"order_state"`
	price			string   	`json:"price"`
	order_line		[]OrderLine	`json:"order_line"`

}

type AllOrders struct {
	allOrders	[]Order `json:"allOrders"`
}