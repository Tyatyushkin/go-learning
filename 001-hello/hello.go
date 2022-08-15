package main

import "fmt"

func Hello(name string) string {
	if name == "" {
		name = "GO"
	}
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("GO"))
}
