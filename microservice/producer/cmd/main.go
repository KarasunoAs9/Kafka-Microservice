package main

import (
	"log"
	"net/http"
)



type Order struct {
	CustomerName string `json:"customer_name"`
	CoffeType string `json:"coffe_type"`
}


func main() {
	http.HandleFunc("/order", placeOrder)
	log.Fatal(http.ListenAndServe(":3000", nil))
}	

func placeOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}


	
}