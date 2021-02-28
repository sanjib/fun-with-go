package main

import "fmt"

// matrix

const (
	rows = 4 * 1024
	cols = 4 * 1024
)

var matrix [rows][cols]byte

func ColumnTraverse() int {
	var counter int
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if matrix[row][col] == 0xff {
				counter++
			}
		}
	}
	return counter
}

func RowTraverse() int {
	var counter int
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0xff {
				counter++
			}
		}
	}
	return counter
}

// linked list

type data struct {
	v byte
	p *data
}

var list *data

// LinkedListTraverse traverses the list linearly.
func LinkedListTraverse() int {
	var counter int
	d := list
	for d != nil {
		if d.v == 0xff {
			counter++
		}
		d = d.p
	}
	return counter
}

// init
func init() {
	var last *data
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			var d data
			// initialize the list
			if list == nil {
				list = &d
			}
			if last != nil {
				last.p = &d
			}
			last = &d

			if row%2 == 0 {
				matrix[row][col] = 0xff
				d.v = 0xff
			}
		}
	}

	var counter int
	d := list
	for d != nil {
		counter++
		d = d.p
	}

	fmt.Println("elements in the link list", counter)
	fmt.Println("elements in the matrix", rows*cols)
}

// main
func main() {
	fmt.Println("list", LinkedListTraverse())
	fmt.Println("by col", ColumnTraverse())
	fmt.Println("by row", RowTraverse())
}
