package main

import "fmt"

func checkType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("The string is %s\n", v)
	case int:
		fmt.Printf("The integer is %d\n", v)
	default:
		fmt.Println("I don't know about that type")
	}
}

func main() {
	checkType("Hello")
}
