package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	// Define the remote host and port range
	host := "192.168.3.14"
	startPort := 1
	endPort := 9999

	start := time.Now()
	// Create a wait group to synchronize the goroutines
	var wg sync.WaitGroup

	// Loop through the port range and start a goroutine for each port
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", host, port)
			conn, err := net.DialTimeout("tcp", address, 10*time.Second)
			if err == nil {
				fmt.Printf("Port %d is open\n", port)
				conn.Close()
			}
		}(port)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	end := time.Now()
	fmt.Printf("Execution time: %v\n", end.Sub(start))
}
