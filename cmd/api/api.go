package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"food_order/cmd/api/routes"

	"github.com/gorilla/mux"
)

func indexAction(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	fmt.Fprint(writer, `["Testing 1 2 3"]`)
}

func getAllOrdersAction(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	//orders := data.GetAllOrders()
	json.NewEncoder(writer).Encode(request)
}

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexAction).Methods("GET")
	router.HandleFunc("/order", getAllOrdersAction).Methods("GET")
	router.HandleFunc("/order", routes.CreateOrderAction).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}