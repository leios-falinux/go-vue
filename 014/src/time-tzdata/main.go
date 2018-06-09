package main

import "fmt"
import "time"
import "github.com/gobuffalo/packr"

func main() {
	box := packr.NewBox("tzdata")

	t, err := box.MustBytes("Seoul")
	if err != nil {
		panic(err)
	}

	l, err := time.LoadLocationFromTZData("Asia/Seoul", t)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	fmt.Println(now.In(l))
}
