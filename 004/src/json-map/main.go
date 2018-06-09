package main

import "fmt"
import "encoding/json"

func main() {
	jsonData := `{  "name": "minsoo",  "team": "sw" }`

	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		// handle error
		panic(err)
	}

	name, ok := data["name"].(string)
	if !ok {
	}
	team, ok := data["team"].(string)
	if !ok {
	}

	fmt.Println(name, team)
}
