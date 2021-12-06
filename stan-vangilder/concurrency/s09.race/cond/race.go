package main

import (
	"fmt"
	"sync"
)

type widgetInventory struct {
	mu          sync.Mutex
	newPurchase *sync.Cond
	quantity    int
}

func main() {
	var wg sync.WaitGroup

	inventory := &widgetInventory{
		quantity: 10,
	}
	inventory.newPurchase = sync.NewCond(&inventory.mu)

	fmt.Println("starting inventory count =", inventory.quantity)

	wg.Add(2)
	go makeSales(inventory, &wg)
	go makePurchases(inventory, &wg)
	wg.Wait()

	fmt.Println("current inventory count =", inventory.quantity)
}

func makeSales(inventory *widgetInventory, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 30_000; i++ {
		inventory.mu.Lock()
		for inventory.quantity-100 < 0 {
			inventory.newPurchase.Wait()
		}
		inventory.quantity -= 100
		fmt.Println("in sales", inventory.quantity)
		inventory.mu.Unlock()
	}
}

func makePurchases(inventory *widgetInventory, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 30_000; i++ {
		inventory.mu.Lock()
		inventory.quantity += 100
		fmt.Println("in purchases", inventory.quantity)
		inventory.newPurchase.Signal()
		inventory.mu.Unlock()
	}
}
