package mod01

func NumInList(list []int, n int) bool {
	for _, v := range list {
		if n == v {
			return true
		}
	}
	return false
}
