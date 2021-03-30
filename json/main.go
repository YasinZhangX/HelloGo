package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Order struct {
	ID          string      `json:"id"`
	Item        []OrderItem `json:"item"`
	Quantity    int         `json:"quantity"`
	TotoalPrice float64     `json:"total_price"`
}

func marshal() {
	o := Order{
		ID: "123456",
		Item: []OrderItem{
			{
				ID:    "item_1",
				Name:  "learn go",
				Price: 15,
			},
			{
				ID:    "item_1",
				Name:  "learn go",
				Price: 15,
			},
		},
		Quantity:    3,
		TotoalPrice: 45,
	}

	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)
}

func unmarshal() {
	s := `{"id":"123456","item":[{"id":"item_1","name":"learn go","price":15},{"id":"item_1","name":"learn go","price":15}],"quantity":3,"total_price":45}`
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", o)
}

func main() {
	marshal()
	unmarshal()
}
