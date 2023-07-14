package arrays

import (
	"fmt"
)

func A1() {
	sum := 1
	for i := 0; i < 20; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
