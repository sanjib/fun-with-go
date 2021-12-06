package main

import (
	"fmt"
	"sync"
)

type widgetInventory struct {
	mu       sync.Mutex
	quantity int
}

func main() {
	var wg sync.WaitGroup

	inventory := &widgetInventory{
		quantity: 1000,
	}

	fmt.Println("starting inventory count =", inventory.quantity)

	wg.Add(2)
	go makeSales(inventory, &wg)
	go makePurchases(inventory, &wg)
	wg.Wait()

	fmt.Println("current inventory count =", inventory.quantity)
}

func makeSales(inventory *widgetInventory, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 300_000; i++ {
		inventory.mu.Lock()
		inventory.quantity--
		inventory.mu.Unlock()
	}
}

func makePurchases(inventory *widgetInventory, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 300_000; i++ {
		inventory.mu.Lock()
		inventory.quantity++
		inventory.mu.Unlock()
	}
}
