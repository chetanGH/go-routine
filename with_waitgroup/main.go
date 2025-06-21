package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/chetanGH/demo/data"
)

func main() {
	startTime := time.Now() // Record the current time
	var totalSales int
	var wg sync.WaitGroup

	wg.Add(len(data.SalesData)) // Set the number of goroutines to wait for

	for _, salesRegion := range data.SalesData {
		// Invoke to calculate Total for Each Region
		go calculateRegionSales(salesRegion, &totalSales, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Printf("Total %d , Time taken to calculate %s \n", totalSales, time.Since(startTime))
}

// Function to Calculate Total Sales per Region
func calculateRegionSales(salesRegion []int, totalSales *int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done when the function returns
	regionTotal := 0
	for _, storeSales := range salesRegion {
		// Calculate region total
		regionTotal += storeSales
		time.Sleep(100 * time.Millisecond)
	}

	// Post Calculation of regionTotal
	*totalSales += regionTotal
}
