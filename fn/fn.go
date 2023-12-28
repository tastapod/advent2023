package fn

import "fmt"

func Reduce[T any](input []T, reduce func(T, T) T) (acc T) {
	if len(input) < 2 {
		panic(fmt.Sprintf("Cannot reduce %v", input))
	}
	acc = reduce(input[0], input[1])
	for i := 2; i < len(input); i++ {
		acc = reduce(acc, input[i])
	}
	return
}
