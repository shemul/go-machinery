package tasks

func Add(nums ...int64) int64 {
	sum := int64(0)
	for _, n := range nums {
		sum += n
	}
	return sum
}
