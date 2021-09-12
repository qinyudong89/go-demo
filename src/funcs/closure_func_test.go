package funcs

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	fib := Fib()
	for i := 0; i < 10; i++ {
		fmt.Print(fib(), " ")
	}
}
