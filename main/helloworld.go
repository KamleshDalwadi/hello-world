package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println(greet("All"))
}

// greet greets saying Hi
func greet(input string) (string, error) {
	return "Hi " + input, nil
}
