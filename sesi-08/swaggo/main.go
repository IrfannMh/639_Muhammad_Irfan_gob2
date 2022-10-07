package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "swaggo/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

type Order struct {
	OrderID      string    `json:"orderId" example:"1"`
	CustomerName string    `json:"customerName" example:"Leo Messi"`
	OrderedAt    time.Time `json:"orderedAt" example:"2019-l1-09T21:21:46+00:00"`
	Items        []Item    `json:"items"`
}

type Item struct {
	ItemId      string `json:"itemId" example:"ABC123"`
	Description string `json:"description" example:"lorem ipsum"`
	Quantity    int    `json:"quantity" example:"1"`
}

var orders []Order
var prevOrderID = 0

// @title Orders API
// @version 1.0
// @description This is a sample service for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	router := mux.NewRouter()

	//create
	router.HandleFunc("/order", createOrder).Methods("POST")

	//get all
	router.HandleFunc("/order", getOrders).Methods("GET")

	//get data by id
	router.HandleFunc("/order/{orderId}", getOrder).Methods("GET")

	//delete data by id
	router.HandleFunc("/order/{orderId}", deleteOrder).Methods("DELETE")

	//update data by id
	router.HandleFunc("/order/{orderId}", updateOrder).Methods("PUT")

	//swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	//call to browser http://localhost:8080/swagger/index.html

	log.Fatal(http.ListenAndServe(":8080", router))
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input paylod
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body Order true "Create order"
// @Success 200 {object} Order
// @Router /orders [post]
func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderID++
	order.OrderID = strconv.Itoa(prevOrderID)
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

// GetOrders godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} Order
// @Router /orders [get]
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)

}

// GetOrderByID godoc
// @Summary Get details orders by id
// @Description Get details order by id
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID"
// @Success 200 {array} Order
// @Router /orders/{orderId} [get]
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderId := params["orderId"]

	for _, order := range orders {
		if order.OrderID == inputOrderId {
			json.NewEncoder(w).Encode(order)
			return
		}
	}
}

// DeleteOrder godoc
// @Summary delete data orders by id
// @Description delete data order by id
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID"
// @Success 204 "No Content"
// @Router /orders/{orderId} [delete]
func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderId := params["orderId"]

	for i, order := range orders {
		if order.OrderID == inputOrderId {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

}

// UpdateOrderByID godoc
// @Summary Update data orders by id
// @Description Update data order by id
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID" of the order
// @Success 200 {array} Order
// @Router /orders/{orderId} [put]
func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderId := params["orderId"]

	for i, order := range orders {
		if order.OrderID == inputOrderId {
			orders = append(orders[:i], orders[i+1:]...)
			var updatedOrder Order
			json.NewDecoder(r.Body).Decode(&updatedOrder)
			orders = append(orders, updatedOrder)
			json.NewEncoder(w).Encode(updatedOrder)
			return
		}
	}

}
