package api

import (
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

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexAction).Methods("GET")
	router.HandleFunc("/order", routes.GetAllOrdersAction).Methods("GET")
	router.HandleFunc("/order", routes.CreateOrderAction).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
