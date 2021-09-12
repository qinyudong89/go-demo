package funcs

import (
	"github.com/qinyudong89/go-demo/src/funcs/panics"
	"testing"
)

func TestDefer(t *testing.T) {
	panics.Calculate(2)
}
