package main

import (
	"fmt"
	"math/rand"
	"time"
	"concurrent-order-processing/order"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	totalOrders := 25
	numWorkers := 5
	orders := make([]order.Order, totalOrders)
	for i := range orders {
		orders[i] = order.Order{ID: i + 1}
	}
	results := order.ProcessOrders(orders, numWorkers)
	fmt.Println("Order Processing Results:")
	for _, res := range results {
		fmt.Printf("Order %2d: %s\n", res.Order.ID, res.Order.State)
	}
}