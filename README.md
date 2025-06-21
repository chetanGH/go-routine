# Go-routine: Concurrency Patterns in Go

## Overview
This project demonstrates various concurrency patterns in Go for calculating total sales across multiple regions. It highlights importance of using goroutines, channels, mutexes, and WaitGroups to manage concurrent execution effectively.

---

## Concurrency Patterns

### 1. Sequential Calculation (No Concurrency)
**File:** `main.go`  
**Description:** This implementation calculates total sales sequentially without concurrency. It serves as a baseline for comparing performance improvements with concurrent approaches.

#### Code Example:
```go
var totalSales int
for _, salesRegion := range data.SalesData {
    calculateRegionSales(salesRegion, &totalSales)
}
```
**note:** Understanding the limitations of sequential processing and identifying opportunities for concurrency.

---

### 2. Goroutines
**File:** `with_goroutine/main.go`  
**Description:** Goroutines are lightweight threads in Go. This implementation uses goroutines to calculate sales concurrently but relies on `time.Sleep` to allow goroutines to complete.

#### Code Example:
```go
for _, salesRegion := range data.SalesData {
    go calculateRegionSales(salesRegion, &totalSales)
}
time.Sleep(3 * time.Second) // Wait for goroutines to complete
```
**note:** Using goroutines for concurrent execution and understanding their lifecycle.

---

### 3. WaitGroup
**File:** `with_waitgroup/main.go`  
**Description:** `sync.WaitGroup` ensures all goroutines complete before proceeding. This approach eliminates the inefficiency of relying on `time.Sleep`.

#### Code Example:
```go
var wg sync.WaitGroup
wg.Add(len(data.SalesData))
for _, salesRegion := range data.SalesData {
    go calculateRegionSales(salesRegion, &totalSales, &wg)
}
wg.Wait() // Wait for all goroutines to finish
```
**note:** Managing goroutines effectively using WaitGroups for synchronization.

---

### 4. Channels
**File:** `with_channel/main.go`  
**Description:** Channels enable communication between goroutines and the main function. This implementation uses channels to collect sales totals from each region.

#### Code Example:
```go
salesCh := make(chan int)
for _, salesRegion := range data.SalesData {
    go calculateRegionSales(salesRegion, salesCh)
}
for i := 0; i < len(data.SalesData); i++ {
    totalSales += <-salesCh // Receive values from the channel
}
close(salesCh)
```
**note:** Leveraging channels for safe and efficient communication between goroutines.

---

### 5. Mutex
**File:** `with_mutex/main.go`  
**Description:** `sync.Mutex` protects shared data (`totalSales`) from race conditions when accessed by multiple goroutines.

#### Code Example:
```go
var mu sync.Mutex
for _, salesRegion := range data.SalesData {
    go calculateRegionSales(salesRegion, &totalSales, &wg, &mu)
}
mu.Lock()
*totalSales += regionTotal
mu.Unlock()
```
**note:** Using mutexes to ensure thread-safe access to shared resources.

---

## Sales Data
The sales data is stored in `data/SalesData.go` as a 2D slice, representing sales figures for five regions: East, West, North, South, and Central.

---

## How to Run
1. Clone the repository.
2. Navigate to the desired concurrency implementation directory.
3. Run the program using:
   ```sh
   go run main.go
   ```

---

## ⭐ Support and Learn Together!
If you found this project helpful or learned something new, consider giving it a ⭐ on GitHub!  
Let's learn and grow together!

---
