package tasks

import "time"

func Add(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	time.Sleep(time.Second * 5)

	return sum, nil
}
