package order

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID    int
	State string // "completed" or "failed"
}

type Result struct {
	Order Order
}

// ProcessOrder simulates processing an order, potentially failing at random
func ProcessOrder(order Order) Result {
	time.Sleep(time.Duration(rand.Intn(50)+10) * time.Millisecond) // simulate order processing time
	if rand.Intn(3) == 0 { // about 1/3 chance to fail
		order.State = "failed"
	} else {
		order.State = "completed"
	}
	return Result{Order: order}
}

// Worker receives jobs on jobsCh and writes results to resultsCh
func Worker(id int, jobsCh <-chan Order, resultsCh chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range jobsCh {
		res := ProcessOrder(order)
		resultsCh <- res
	}
}

// ProcessOrders concurrently processes orders with a fixed number of workers
func ProcessOrders(orders []Order, numWorkers int) []Result {
	jobsCh := make(chan Order)
	resultsCh := make(chan Result)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(i, jobsCh, resultsCh, &wg)
	}

	// Feed orders to jobsCh (in a separate goroutine)
	go func() {
		for _, order := range orders {
			jobsCh <- order
		}
		close(jobsCh)
	}()

	// Collect results
	var results []Result
	go func() {
		wg.Wait()
		close(resultsCh)
	}()
	for res := range resultsCh {
		results = append(results, res)
	}

	return results
}
