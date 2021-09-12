package panics

import "testing"

func TestPanic(t *testing.T) {
	names := []string{"Java", "PHP", "C", "C++"}
	for i := 0; i < len(names); i++ {
		Language(names[i])
	}
}
