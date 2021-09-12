package panics

import "fmt"

func Calculate(num int) error {
	for i := 0; i < 5; i++ {
		divisor := num - i
		if divisor == 0 {
			panic("0 cannot be a divisor")
		}
		result := i / divisor
		fmt.Println(result)
	}
	return nil
}
