package mod01

func Sum(list []int) int {
	var sum int
	for _, v := range list {
		sum += v
	}
	return sum
}

func Sum2(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + Sum2(list[1:])
}
