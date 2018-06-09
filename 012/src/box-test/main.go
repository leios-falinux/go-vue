package main

import (
	"fmt"
	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("tzdata")

	s := box.String("abc.txt")
	fmt.Println(s)
}
