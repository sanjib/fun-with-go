package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type widgetInventory struct {
	quantity int32
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
		atomic.AddInt32(&inventory.quantity, 1)
	}
}

func makePurchases(inventory *widgetInventory, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 300_000; i++ {
		atomic.AddInt32(&inventory.quantity, -1)
	}
}
