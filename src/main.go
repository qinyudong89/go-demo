package main

import (
	"fmt"
	"github.com/qinyudong89/go-demo/src/funcs"
)

func main() {
	//result :=funcs.Hypot(2.0,4.0)
	//fmt.Println("result: %d",result)
	f := funcs.Squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
