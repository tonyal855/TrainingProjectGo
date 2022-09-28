package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `example:"STATUS_OK`
	Status  int    `example:"200"`
	Payload interface{}
}

type Order struct {
}

// @title Orders api
// @description Sample Api Order
// @version v1.0
// @termOfService http://swagger.io/terms/
// @host localhost:4000/docs
func main() {

	http.HandleFunc("/orders", getOrders)
	http.ListenAndServe(":4000", nil)

}

// GetOrders godoc
// @Summary Get All order
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {array} Response
// @Router /orders [get]
func getOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{
		Message: "OK",
		Status:  200,
		Payload: nil,
	})
}
