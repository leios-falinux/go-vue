package main

import "fmt"
import "encoding/json"
import "io/ioutil"
import "github.com/imdario/mergo"

type ConfigData struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	RootDir string `json:"rootdir"`
}

func main() {
	var config ConfigData

	default_config_content := `{
	        "address": "127.0.0.1",
	        "port": 3000,
	        "rootdir": "/home/leios/falinux"
	}`

	if err := json.Unmarshal([]byte(default_config_content), &config); err != nil {
		panic(err)
	}

	// override config
	if raw, err := ioutil.ReadFile("config.json"); err == nil {
		var new_config ConfigData

		if err := json.Unmarshal(raw, &new_config); err != nil {
			panic(err)
		}

		if err := mergo.Merge(&config, new_config, mergo.WithOverride); err != nil {
			panic(err)
		}
	}
	fmt.Printf("config: [%+v]\n", config)
}
