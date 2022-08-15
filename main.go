package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rest-api/handler"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/allgroceries", handler.AllGroceries)                        // ----> To request all groceries
	r.HandleFunc("/groceries/{name}/{quantity}", handler.SingleGrocery)        // ----> To request a specific grocery
	r.HandleFunc("/groceries", handler.GroceriesToBuy).Methods("POST")         // ----> To add  new grocery to buy
	r.HandleFunc("/groceries/{name}", handler.UpdateGrocery).Methods("PUT")    // ----> To update a grocery
	r.HandleFunc("/groceries/{name}", handler.DeleteGrocery).Methods("DELETE") // ----> Delete a grocery
	log.Fatal(http.ListenAndServe(":10000", r))
}
