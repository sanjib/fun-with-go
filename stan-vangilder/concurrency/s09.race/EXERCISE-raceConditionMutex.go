//	USING CHANNELS INSTEAD OF SHARING VARIABLES (widgetInventory in this case)
//	This file manages inventory transactions.  We've mentioned that channels are the preferred method for sharing data.
//	Mutexes are not difficult, but your mission is to modify this code to avoid variable sharing by using channels.
//	Before you start, close all other open files in your IDE.
//	When operating correctly, this program should report 1000 as both starting and ending inventories.
//	1) Run this code now and verify correct operation.  You should see starting and ending inventories as 1000.
//	2) Modify the code to avoid sharing widgetInventory in the goroutines and instead use channels for communication.
//	3) If you need hints, look at the bottom of this file, or open the SOLUTION file. Use hints only as needed.
//	4) Verify proper operation as above after your modifications.
//	5) Make other changes and experiment as you like. Take your time and be creative!
//	6) Share any interesting discoveries.

package main

import (
	"fmt"
)

var (
	widgetInventory int32 = 1000
)

func main() {
	chSales := make(chan int32)
	chPurchases := make(chan int32)

	fmt.Println("Starting inventory count = ", widgetInventory)
	go makeSales(chSales)
	go newPurchases(chPurchases)

	for i := 0; i < 6000; i++ {
		select {
		case val := <-chSales:
			widgetInventory -= val
		case val := <-chPurchases:
			widgetInventory += val
		}
	}

	fmt.Println("Ending inventory count = ", widgetInventory)
}

func makeSales(ch chan int32) {
	for i := 0; i < 3000; i++ {
		ch <- 100
	}
}

func newPurchases(ch chan int32) {
	for i := 0; i < 3000; i++ {
		ch <- 100
	}
}

//
// HINTS BELOW
//
//
//
//
//
//	There are many ways to approach this exercise.  This is just ONE way:
// 	Consider:
//	1) Create a large buffered channel to contain all sales and purchase transactions, +100 or -100.
//	2) Load each transaction into the channel with makeSales (-100) and newPurchases (+100).
// 	3) Read all the transactions from the channel and apply them to widgetInventory in func main().
//  4) Open SOLUTION file as needed.
