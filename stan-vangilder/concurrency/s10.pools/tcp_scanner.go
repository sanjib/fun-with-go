package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		fmt.Printf("%d ", p)
		addr := fmt.Sprintf(":%d", p)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)

	// initiate workers
	fmt.Println("checking ports: ")
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	// assign jobs
	go func() {
		for i := 0; i < 1024; i++ {
			ports <- i
		}
	}()

	// get results
	var openPorts []int
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("\nok: %d", port)
	}
}
