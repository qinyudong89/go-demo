package funcs

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	err := Read("https://www.baidu.com44/")
	fmt.Println(err)
}
