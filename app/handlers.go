package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name     string `json:"full_name" xml:"name"`
	City     string `json:"city" xml:"city"`
	Postcode string `json:"postcode" xml:"postcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!\n")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Sandeep Bomma", City: "Belfast", Postcode: "BT39FJ"},
		{Name: "Kalyani Yele", City: "Deverakonda", Postcode: "BT16BB"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application.xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
