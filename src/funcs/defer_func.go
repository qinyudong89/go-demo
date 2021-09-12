package funcs

import (
	"fmt"
	"net/http"
)

func Read(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	//Here defer should be written after err judgment(If the resource is not successful, there is no need to release the resource)
	defer resp.Body.Close()
	fmt.Print(resp)
	return nil
}

//If there are multiple defer functions, the calling sequence is similar to the stack
func execSort() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
}
