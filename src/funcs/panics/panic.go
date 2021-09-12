package panics

import "fmt"

func Language(name string) {
	switch name {
	case "C":
		fmt.Println(name)
		break
	case "C++":
		fmt.Println(name)
		break
	case "Java":
		fmt.Println(name)
		break
	case "Golang":
		fmt.Println(name)
		break
	default:
		panic(fmt.Sprintf("invalid name %q", name))
	}
}
