package funcs

import (
	"fmt"
	"net/http"
)

var conext string = ""

func Read(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Print(resp)
	return nil
}

func Write(text string) {
	conext = text
}
