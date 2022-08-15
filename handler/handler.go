package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rest-api/model"
)

var groceries = []model.Grocery{
	{Name: "Almod Milk", Quantity: 2},
	{Name: "Apple", Quantity: 6},
}

func AllGroceries(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: returnAllGroceries")
	json.NewEncoder(w).Encode(groceries)
}

func SingleGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	quantity := vars["quantity"]
	fmt.Println(quantity)

	for _, grocery := range groceries {
		if grocery.Name == name {
			json.NewEncoder(w).Encode(grocery)
		}
	}
}

func GroceriesToBuy(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var grocery model.Grocery
	json.Unmarshal(reqBody, &grocery)
	groceries = append(groceries, grocery)

	json.NewEncoder(w).Encode(groceries)

}

func UpdateGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	for index, grocery := range groceries {
		if grocery.Name == name {
			groceries = append(groceries[:index], groceries[index+1:]...)

			var updateGrocery model.Grocery

			json.NewDecoder(r.Body).Decode(&updateGrocery)
			groceries = append(groceries, updateGrocery)
			fmt.Println("Endpoint hit: UpdateGroceries")
			json.NewEncoder(w).Encode(updateGrocery)
			return
		}
	}

}

func DeleteGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	for index, grocery := range groceries {
		if grocery.Name == name {
			groceries = append(groceries[:index], groceries[index+1:]...)
		}
	}

}
