package order

import (
	"testing"
)

func TestProcessOrders(t *testing.T) {
	totalOrders := 25
	numWorkers := 4
	orders := make([]Order, totalOrders)
	for i := range orders {
		orders[i] = Order{ID: i + 1}
	}
	results := ProcessOrders(orders, numWorkers)
	if len(results) != totalOrders {
		t.Errorf("expected %d results, got %d", totalOrders, len(results))
	}
	seen := make(map[int]bool)
	completed, failed := 0, 0
	for _, r := range results {
		if seen[r.Order.ID] {
			t.Errorf("order ID %d processed more than once", r.Order.ID)
		}
		seen[r.Order.ID] = true
		if r.Order.State == "completed" {
			completed++
		} else if r.Order.State == "failed" {
			failed++
		} else {
			t.Errorf("unexpected state: %s", r.Order.State)
		}
	}
	if completed+failed != totalOrders {
		t.Errorf("some orders not processed: completed=%d, failed=%d", completed, failed)
	}
}
