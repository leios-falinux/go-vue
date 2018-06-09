package main

import "fmt"

var version string
var build string

func main() {
	fmt.Printf("App Version: %s\n", version)
	fmt.Printf("build mode: %s\n", build)

	if build == "release" {
		fmt.Println("  release actions...")
	} else {
		fmt.Println("  debug actions...")
	}
}
