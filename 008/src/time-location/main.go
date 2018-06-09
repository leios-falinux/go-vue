package main

import "fmt"
import "time"

func main() {
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}
	now := time.Now()
	fmt.Println(now.In(location))
}
