package main

import "fmt"
import "encoding/json"

func main() {
	jsonData := `{  "name": "minsoo",  "team": "sw" }`

	type Member struct {
		Name string `json:"name"`
		Team string `json:"team"`
	}

	m := Member{}
	err := json.Unmarshal([]byte(jsonData), &m)
	if err != nil {
		// handle error
	}

	fmt.Println(m.Name, m.Team)
}
