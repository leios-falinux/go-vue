package main

import "fmt"
import "time"
import "io/ioutil"

func main() {
  t, err := ioutil.ReadFile("Seoul")
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
