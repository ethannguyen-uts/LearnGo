// fib_stress.go
package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"
)

// fibInefficient computes Fibonacci using naive recursion,
// launching two goroutines for n-1 and n-2 on every call.
// It also performs a small busy loop to increase CPU usage.
// Returns *big.Int to avoid overflow for moderately large n.
func fibInefficient(n int) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}

	// Channel to receive two results
	ch := make(chan *big.Int, 2)

	// Spawn goroutines for both halves (this explodes goroutine count)
	go func() {
		ch <- fibInefficient(n - 1)
	}()
	go func() {
		ch <- fibInefficient(n - 2)
	}()

	// Receive results (order not important)
	a := <-ch
	b := <-ch

	// Busy work to burn CPU (adjust the loop count to increase/decrease load)
	// Keep this moderate so it stresses CPU rather than crashing the system.
	for i := 0; i < 5_000_000; i++ {
		_ = i * i
	}

	sum := new(big.Int).Add(a, b)
	return sum
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run fib_stress.go <n>")
		fmt.Println("Example: go run fib_stress.go 36")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		fmt.Println("Please provide a non-negative integer for n.")
		return
	}

	fmt.Printf("Starting extremely inefficient Fibonacci for n=%d\n", n)
	start := time.Now()
	result := fibInefficient(n)
	elapsed := time.Since(start)

	fmt.Printf("Fib(%d) = %s\n", n, result.String())
	fmt.Printf("Elapsed: %s\n", elapsed)
}
