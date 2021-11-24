package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
)

func main() {
	numsFlag := flag.String("nums", "", "Numbers, usage: max_value -nums=\"16 8 42 4 23 15\"")
	flag.Parse()

	var nums []int
	for _, numFlag := range strings.Split(*numsFlag, " ") {
		num, err := strconv.Atoi(strings.TrimSpace(numFlag))
		if err != nil {
			log.Println(err)
		}
		nums = append(nums, num)
	}

	if len(nums) == 0 {
		log.Println("Max:", nil)
		return
	}

	if len(nums) == 1 {
		log.Println("Max:", nums[0])
		return
	}

	max := nums[0]

	//for i := 1; i < len(nums); i++ {
	//	if nums[i] > max {
	//		max = nums[i]
	//	}
	//}
	for _, val := range nums[1:] {
		if val > max {
			max = val
		}
	}

	log.Println("Max:", max)
}
