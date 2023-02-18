package main

import (
	"fmt"

	gt "github.com/bas24/googletranslatefree"
)

func main() {
	const text string = "hello world"
	result, _ := gt.Translate(text, "en", "my")
	fmt.Println(result)
}
